package scheduler

import (
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/entities"
	myconfig "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-worker-service/config"
	"sync"
)

//任务调度
type Scheduler struct {
	taskEventChan chan *entities.TaskEvent // etcd任务事件队列
}

var (
	SgtScheduler *Scheduler
	once sync.Once
)

// 初始化调度器单例
func InitSchedulerSingleton() *Scheduler  {
	once.Do(func() {
		SgtScheduler = &Scheduler{
			taskEventChan: make(chan *entities.TaskEvent, myconfig.ConfigSingleton.TaskEventCapacity),
		}
	})
	return SgtScheduler
}

// 初始化调度表
func InitScheduler() (err error)  {
	go SgtScheduler.ScheduleLoop()
	return
}

// 调度协程
func (s *Scheduler) ScheduleLoop()  {

}