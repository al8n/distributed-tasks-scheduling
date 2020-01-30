package distribution

import (
	myconfig "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/config"
	"go.etcd.io/etcd/clientv3"
	"sync"
	"time"
)

type EtcdInterface interface {

}

type EtcdDB struct {
	client *clientv3.Client
	kv clientv3.KV
	lease clientv3.Lease
}

var (
	EtcdSingleton *EtcdDB
	once sync.Once
)

func newEtcdSingleton(client *clientv3.Client, kv clientv3.KV, lease clientv3.Lease) *EtcdDB  {
	once.Do(func() {
		EtcdDBSingleton = &EtcdDB{
			client: client,
			kv: kv,
			lease: lease,
		}
	})

	return EtcdSingleton
}

// 初始化etcd连接
func InitEtcd() (err error)  {
	var (
		config clientv3.Config
		client *clientv3.Client
		lease clientv3.Lease
		kv clientv3.KV
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

	// 得到KV和lease的API子集
	kv = clientv3.NewKV(client)
	lease = clientv3.NewLease(client)

	// 赋值单例
	newEtcdSingleton(client, kv, lease)
	return
}

