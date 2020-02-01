package distribution

import (
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/entities"
	"os/exec"
	"sync"
	"time"
)

type executor interface {
	ExecuteTask(info *entities.TaskExecuteInfo)
}

type Executor struct {

}

var (
	SgtExecutor *Executor
	onceExecutor sync.Once
)

func InitExecutor() *Executor  {
	onceExecutor.Do(func() {
		SgtExecutor = &Executor{}
	})
	return SgtExecutor
}

// 执行一个任务
func (e *Executor) ExecuteTask(info *entities.TaskExecuteInfo)  {
	go func() {
		var (
			cmd *exec.Cmd
			err error
			output []byte
			rst *entities.TaskExecuteResult
			lock *Lock
		)

		// 任务结果
		rst  = &entities.TaskExecuteResult{
			ExecuteInfo: info,
			Output: make([]byte, 0),
		}

		// 初始化锁
		lock = SgtEtcd.CreateTaskLock(info.Task.Name)

		// 记录任务开始时间
		rst.StartTime = time.Now()

		err = lock.TryLock()
		// 释放锁
		defer lock.UnLock()

		if err != nil { // 上锁失败
			rst.Err = err
			rst.EndTime = time.Now()
		} else {
			// 上锁成功后，重置任务启动时间
			rst.StartTime = time.Now()

			// 执行shell命令
			cmd = exec.CommandContext(info.Context, "/bin/bash", "-c", info.Task.Command)

			// 执行并捕获输出
			output, err = cmd.CombinedOutput()

			// 记录任务结束时间
			rst.EndTime = time.Now()
			rst.Output = output
			rst.Err = err

		}
		// 把任务执行完成后， 把执行的结果返回给Scheduler，
		// Scheduler会从executing table中 删除掉执行记录
		SgtScheduler.PushTaskResultBack(rst)
	}()
}
