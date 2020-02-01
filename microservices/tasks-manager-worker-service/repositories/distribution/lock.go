package distribution

import (
	"context"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/conf"
	myutils "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/utils"
	"github.com/coreos/etcd/clientv3"
	"log"
)

type lock interface {
	TryLock()  (err error)
	UnLock()
}

// 分布式锁(TXN事务)
type Lock struct {
	kv clientv3.KV
	lease clientv3.Lease

	taskName string  //任务名
	cancelFunc context.CancelFunc // 用于终止自动续租
	leaseID clientv3.LeaseID
	isLocked bool // 是否上锁成功
}

// 初始化一把锁
func InitTaskLock(taskName string, kv clientv3.KV, lease clientv3.Lease) (lock *Lock) {
	lock = &Lock{
		kv: kv,
		lease: lease,
		taskName: taskName,
	}
	return
}

// 尝试上锁
func (l *Lock) TryLock() (err error) {
	var (
		leaseGrantResp *clientv3.LeaseGrantResponse
		ctx context.Context
		cancelFunc context.CancelFunc
		leaseID clientv3.LeaseID
		keepRespChan <- chan *clientv3.LeaseKeepAliveResponse
		txn clientv3.Txn
		lockKey string
		txnResp *clientv3.TxnResponse
	)

	// 1. 创建租约
	if leaseGrantResp, err = l.lease.Grant(context.TODO(), 5); err != nil {
		return
	}

	// context 用于取消自动续租
	ctx, cancelFunc = context.WithCancel(context.TODO())

	// 租约ID
	leaseID = leaseGrantResp.ID

	// 声明一个释放租约的函数
	f := func() {
		cancelFunc()
		l.lease.Revoke(context.TODO(), leaseID)
	}

	// 2. 自动续租
	if keepRespChan, err = l.lease.KeepAlive(ctx, leaseID); err != nil {
		f() // 释放租约
		return
	}

	// 处理续租应答的协程
	go func() {
		var (
			keepResp *clientv3.LeaseKeepAliveResponse
		)
		for {
			select {
			case keepResp = <- keepRespChan: // 自动续租的应答
				if keepResp == nil {
					break
				}
			}
		}
	}()

	// 3. 创建事务
	txn = l.kv.Txn(context.TODO())
	// 锁路径
	lockKey = conf.TASK_LOCK_DIR + l.taskName

	// 4. 事务抢锁
	txn.If(clientv3.Compare(clientv3.CreateRevision(lockKey), "=", 0)).
		Then(clientv3.OpPut(lockKey, "", clientv3.WithLease(leaseID))).
		Else(clientv3.OpGet(lockKey))

	// 提交任务
	if txnResp, err = txn.Commit(); err != nil {
		f()
	}
	// 5. 成功返回，失败释放租约
	if !txnResp.Succeeded { // 锁被占用
		err = myutils.LockOccupiedError
		f()
	}

	// 抢锁成功
	log.Println("successful\n")
	l.leaseID = leaseID
	l.cancelFunc = cancelFunc
	l.isLocked = true
	return
}

func (l *Lock) UnLock()  {
	if l.isLocked {
		l.cancelFunc() // 取消自动续租的协程
		l.lease.Revoke(context.TODO(), l.leaseID) // 释放租约
	}
}