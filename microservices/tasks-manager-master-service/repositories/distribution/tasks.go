package distribution

import (
	"context"
	"encoding/json"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/conf"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/entities"
	myconfig "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/config"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"time"
)

func (e *EtcdDB) SaveOneTask(task *entities.Task) (oldTask *entities.Task, err error) {
	// 把任务保存到/cron/tasks/任务名 -> json
	var (
		taskKey     string
		taskValue   []byte
		putResp     *clientv3.PutResponse
		oldTasksObj entities.Task
		ctx         context.Context
		cancelFunc  context.CancelFunc
	)

	// ectd 保存的key
	taskKey = conf.TASK_SAVE_DIR + task.Name

	// 任务信息json
	if taskValue, err = json.Marshal(task); err != nil {
		return
	}

	ctx, cancelFunc = context.WithTimeout(context.Background(), time.Duration(myconfig.ConfigSingleton.EtcdTimeout) * time.Millisecond)
	defer cancelFunc()

	// 保存到etcd
	if putResp, err = e.kv.Put(ctx, taskKey, string(taskValue), clientv3.WithPrevKV()); err != nil {
		return
	}

	// 如果是更新，那么返回旧值
	if putResp.PrevKv != nil {
		// 对旧值做一个反序列化
		if err = json.Unmarshal(putResp.PrevKv.Value, &oldTasksObj); err != nil {
			err = nil
			return
		}
		oldTask = &oldTasksObj
	}

	return
}

func (e *EtcdDB) DeleteOneTask(name string) (oldTask *entities.Task, err error) {
	var (
		taskKey     string
		delResp     *clientv3.DeleteResponse
		oldTasksObj entities.Task
		ctx         context.Context
		cancelFunc  context.CancelFunc
	)

	// etcd中保存任务的key
	taskKey = conf.TASK_SAVE_DIR + name

	ctx, cancelFunc = context.WithTimeout(context.Background(), time.Duration(myconfig.ConfigSingleton.EtcdTimeout) * time.Millisecond)
	defer cancelFunc()

	// 从etcd中删除它
	if delResp, err = e.kv.Delete(ctx, taskKey, clientv3.WithPrevKV()); err != nil {
		return
	}

	// 返回被删除的任务信息
	if len(delResp.PrevKvs) != 0 {
		// 解析旧值
		if err = json.Unmarshal(delResp.PrevKvs[0].Value, &oldTasksObj); err != nil {
			err = nil
			return
		}
		oldTask   = &oldTasksObj
	}
	return
}

func (e *EtcdDB) GetOneTask(name string) (task *entities.Task, err error) {
	var (
		taskKey string
		getResp *clientv3.GetResponse
		kvPair  *mvccpb.KeyValue
		ctx context.Context
		cancelFunc context.CancelFunc
	)

	taskKey = conf.TASK_SAVE_DIR + name

	ctx, cancelFunc = context.WithTimeout(context.Background(), time.Duration(myconfig.ConfigSingleton.EtcdTimeout) * time.Millisecond)
	defer cancelFunc()

	if getResp, err = e.kv.Get(ctx, taskKey); err != nil {
		return
	}

	kvPair = getResp.Kvs[0]
	task = &entities.Task{}
	if err = json.Unmarshal(kvPair.Value, task); err != nil {
		return task, err
	}
	return
}

func (e *EtcdDB) GetAllTasks() (tasks []*entities.Task, err error) {
	var (
		dirKey string
		getResp *clientv3.GetResponse
		kvPair  *mvccpb.KeyValue
		task *entities.Task
		ctx context.Context
		cancelFunc context.CancelFunc
	)


	// 任务保存的目录
	dirKey = conf.TASK_SAVE_DIR

	ctx, cancelFunc = context.WithTimeout(context.Background(), time.Duration(myconfig.ConfigSingleton.EtcdTimeout) * time.Millisecond)
	defer cancelFunc()

	// 获取目录下所有任务信息
	if getResp, err = e.kv.Get(ctx, dirKey, clientv3.WithPrefix()); err != nil {
		return
	}

	// 初始化数组空间
	tasks = make([]*entities.Task, 0)

	// 遍历所有任务，进行反序列化
	for _, kvPair = range getResp.Kvs {
		task = &entities.Task{}
		if err = json.Unmarshal(kvPair.Value, task); err != nil {
			err = nil
			continue
		}
		tasks = append(tasks, task)
	}
	return
}

func (e *EtcdDB) KillTask(name string) (err error)  {
	var (
		killerKey string
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseId clientv3.LeaseID
		ctx context.Context
		cancelFunc context.CancelFunc
	)

	// 通知worker杀死对应任务
	killerKey = conf.TASK_KILLER_DIR + name

	ctx, cancelFunc = context.WithTimeout(context.Background(), time.Duration(myconfig.ConfigSingleton.EtcdTimeout) * time.Millisecond)
	defer cancelFunc()

	// 让worker监听到一次put操作, 创建一个租约让其稍后自动过期即可
	if leaseGrantResp, err = e.lease.Grant(ctx, 1); err != nil {
		return
	}

	// 租约ID
	leaseId = leaseGrantResp.ID

	// 设置killer标记
	if _, err = e.kv.Put(context.TODO(), killerKey, "", clientv3.WithLease(leaseId)); err != nil {
		return
	}
	return
}