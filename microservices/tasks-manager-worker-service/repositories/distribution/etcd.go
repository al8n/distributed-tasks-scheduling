package distribution

import (
	"context"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/conf"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/entities"
	myconfig "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-worker-service/config"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"log"
	"sync"
	"time"
)

type etcd interface {
	WatchTasks() (err error)
	WatchKiller()
	CreateTaskLock(taskName  string) (lock *Lock)
}

type EtcdDB struct {
	client *clientv3.Client
	kv clientv3.KV
	lease clientv3.Lease
	watcher clientv3.Watcher
}

var (
	SgtEtcd *EtcdDB
	once sync.Once
)

func newEtcdSingleton(etcd *EtcdDB) *EtcdDB  {
	once.Do(func() {
		SgtEtcd = etcd
	})

	return SgtEtcd
}

// 初始化etcd连接
func InitEtcd() (err error)  {
	var (
		config clientv3.Config
		client *clientv3.Client
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
	etcd := &EtcdDB{
		client:  client,
		kv:      clientv3.NewKV(client),
		lease:   clientv3.NewLease(client),
		watcher: clientv3.NewWatcher(client),
	}


	// 赋值单例
	newEtcdSingleton(etcd)

	// 启动任务监听
	SgtEtcd.WatchTasks()

	// 启动监听killer
	SgtEtcd.WatchKiller()
	return
}

// 监听强杀任务通知
func (e *EtcdDB) WatchKiller() {
	var (
		wc clientv3.WatchChan
		wr clientv3.WatchResponse
		we *clientv3.Event
		te *entities.TaskEvent
		taskName string
		task *entities.Task
	)
	go func() { // 监听协程
		// 监听/cron/killer/目录的后续变化
		wc = e.watcher.Watch(context.TODO(), conf.TASK_KILLER_DIR, clientv3.WithPrefix())

		// 处理监听事件
		for wr = range wc {
			for _, we = range wr.Events {
				switch we.Type {
				case mvccpb.PUT: // 杀死任务事件
					taskName = entities.ExtractKillerName(string(we.Kv.Key))
					task = &entities.Task{Name: taskName}
					te = entities.BuildTaskEvent(entities.KILL, task)
					// 变化推送给scheduler
					SgtScheduler.PushTaskEvent(te)
				case mvccpb.DELETE: // killer标记过期被自动删除
				}


			}
		}
	}()
}

// 监听任务变化
func (e *EtcdDB) WatchTasks() (err error) {
	var (
		getResp *clientv3.GetResponse
		kvpair *mvccpb.KeyValue
		task *entities.Task
		watchStartRevision int64
		wc clientv3.WatchChan
		wr clientv3.WatchResponse
		we *clientv3.Event
		te *entities.TaskEvent
	)

	// 1. get一下/cron/tasks/目录下的所有任务，并且获知当前集群的revision
	if getResp, err = e.kv.Get(context.TODO(), conf.TASK_SAVE_DIR, clientv3.WithPrefix()); err != nil {
		return
	}

	// 当前有哪些任务
	for _, kvpair = range getResp.Kvs {
		//反序列化
		if task, err = entities.UnMarshallTask(kvpair.Value); err == nil {

			te = entities.BuildTaskEvent(entities.SAVE, task)
			//同步给调度协程scheduler
			SgtScheduler.PushTaskEvent(te)
		}
	}

	// 2，从该revision向后监听变化事件
	go func() { // 监听协程
		// 从GET时刻的后续版本开始
		watchStartRevision = getResp.Header.Revision + 1
		// 监听/cron/tasks/目录的后续变化
		wc = e.watcher.Watch(context.TODO(), conf.TASK_SAVE_DIR, clientv3.WithRev(watchStartRevision), clientv3.WithPrefix())

		// 处理监听事件
		for wr = range wc {
			for _, we = range wr.Events {
				switch we.Type {
				case mvccpb.PUT:
					if task, err = entities.UnMarshallTask(we.Kv.Value); err != nil {
						log.Fatalln(err)
					}
					// 构建一个Event
					te = entities.BuildTaskEvent(entities.SAVE, task)
				case mvccpb.DELETE:
					taskname := entities.ExtractTaskName(string(we.Kv.Key))

					task = &entities.Task{Name: taskname}
					// 构建删除事件
					te = entities.BuildTaskEvent(entities.DELETE, task)
				}

				// 变化推送给scheduler
				SgtScheduler.PushTaskEvent(te)
			}
		}
	}()
	return
}

// 创建任务执行锁
func (e *EtcdDB) CreateTaskLock(taskName  string) (lock *Lock) {
	lock = InitTaskLock(taskName, e.kv, e.lease)
	return
}
