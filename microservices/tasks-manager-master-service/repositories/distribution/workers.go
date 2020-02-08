package distribution

import (
	"context"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/conf"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/entities"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

func (e *EtcdDB) GetAllWorkers() (workers []string, err error) {
	var (
		getResp *clientv3.GetResponse
		kv *mvccpb.KeyValue
	)
	// 初始化数组
	workers = make([]string, 0)

	// 获取目录下所有kv
	if getResp, err = e.kv.Get(context.TODO(), conf.TASK_WORKER_DIR, clientv3.WithPrefix()); err != nil {
		return
	}

	// 解析每个节点的IP
	for _, kv = range getResp.Kvs {
		// kv.Key
		workerIP := entities.ExtractWorkerIP(string(kv.Key))
		workers = append(workers, workerIP)
	}

	return
}
