package distribution

import (
	"context"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/conf"
	myconfig "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-worker-service/config"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/utils"
	"github.com/coreos/etcd/clientv3"
	"net"
	"sync"
	"time"
)

// 注册节点到etcd /cron/worker/ip
type Register struct {
	client *clientv3.Client
	kv clientv3.KV
	lease clientv3.Lease

	ip string // 本机ip
}

var (
	SingleRegister *Register
	onceRegister sync.Once
)

// 获取本机网卡IP
func getIP() (ip string, err error) {
	var (
		addrs []net.Addr
		addr net.Addr
		ipNet *net.IPNet
		isIPNet bool
	)

	// 获取所有网卡
	if addrs, err = net.InterfaceAddrs(); err != nil {
		return
	}

	// 取消第一个非lo的网卡ip
	for _, addr = range addrs {
		// 这个网络地址是IP地址：ipv4，ipv6
		if ipNet, isIPNet = addr.(*net.IPNet); isIPNet && !ipNet.IP.IsLoopback() {
			// 跳过IPv6
			if ipNet.IP.To4() != nil {
				ip = ipNet.IP.String()
				return
			}
		}
	}

	err = utils.ErrorNoLocalIPFound
	return
}

func InitRegister()  {
	newRegisterSingleton()
	go SingleRegister.keepOnline()
}

func newRegisterSingleton() *Register  {
	onceRegister.Do(func() {
		var (
			config clientv3.Config
			client *clientv3.Client
			err error
			localIP string
		)

		// 初始化配置
		config = clientv3.Config{
			Endpoints: myconfig.ConfigSingleton.EtcdEndpoints,
			DialTimeout: time.Duration(myconfig.ConfigSingleton.EtcdTimeout) * time.Millisecond,
		}

		// 建立连接
		if client, err = clientv3.New(config); err != nil {
			return
		}

		// 本机IP
		if localIP, err = getIP(); err != nil {
			return
		}


		SingleRegister = &Register{
			client:  client,
			kv:      clientv3.NewKV(client),
			lease:   clientv3.NewLease(client),
			ip: localIP,
		}
	})
	return SingleRegister
}

// 注册到/cron/workers/ip 并自动续租
func (r *Register) keepOnline()  {
	var (
		regKey string
		leaseGrantResp *clientv3.LeaseGrantResponse
		err error
		keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
		keepAliveResp *clientv3.LeaseKeepAliveResponse
		ctx context.Context
		cancelFunc context.CancelFunc
	)

	for {
		// 注册路径
		regKey = conf.TASK_WORKER_DIR + r.ip

		// 创建租约
		if leaseGrantResp, err = r.lease.Grant(context.TODO(), 10); err != nil {
			time.Sleep(1 *time.Second)
		}

		// 自动续租
		if keepAliveChan, err = r.lease.KeepAlive(context.TODO(), leaseGrantResp.ID); err != nil {
			time.Sleep(1 *time.Second)
		}

		ctx, cancelFunc = context.WithCancel(context.TODO())

		// 注册到etcd
		if _, err = r.kv.Put(ctx, regKey, "", clientv3.WithLease(leaseGrantResp.ID)); err != nil {
			time.Sleep(1 *time.Second)
			cancelFunc()
		}

		// 处理续租应答
		for {
			select {
			case keepAliveResp = <- keepAliveChan:
				if keepAliveResp == nil {
					time.Sleep(1 *time.Second)
					cancelFunc()
				}
			}
		}
	}
}