package scheduler

import (
	"fmt"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/entities"
	myconfig "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-worker-service/config"
	"sync"
	"time"
)

//任务调度
type Scheduler struct {
	taskEventChan chan *entities.TaskEvent // etcd任务事件队列
	taskPlanTable map[string]*entities.TaskSchedulePlan // 任务调度计划表
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
			taskPlanTable: make(map[string]*entities.TaskSchedulePlan),
		}
	})
	return SgtScheduler
}

// 初始化调度表
func InitScheduler() (err error)  {
	go SgtScheduler.ScheduleLoop()
	return
}

// 重新计算任务调度状态
func (s *Scheduler) TrySchedule() (scheduleAfter time.Duration) {
	var (
		tsp *entities.TaskSchedulePlan
		near *time.Time
	)

	// 如果任务表为空话，随便睡眠多久
	if len(s.taskPlanTable) == 0 {
		scheduleAfter = 1 * time.Second
		return
	}

	// 获取当前时间
	now := time.Now()

	// 1遍历所有任务
	for _, tsp = range s.taskPlanTable {
		if tsp.NextTime.Before(now) || tsp.NextTime.Equal(now) {
			// TODO: 尝试执行任务
			fmt.Println("执行任务", tsp.Task.Name)
			tsp.NextTime = tsp.Expression.Next(now) // 更新下次执行时间
		}

		// 统计最近你一个要过期的任务时间
		if near == nil || tsp.NextTime.Before(*near) {
			near = &tsp.NextTime
		}
	}
	// 下次调度间隔 （最近要执行的任务调度时间-当前时间）
	if near != nil {
		scheduleAfter = (*near).Sub(now)
	}

	return
}

// 调度协程
func (s *Scheduler) ScheduleLoop()  {
	var (
		taskEvent *entities.TaskEvent
		scheduleAfter time.Duration
		scheduleTimer *time.Timer
	)

	// 初始化一次
	scheduleAfter = s.TrySchedule()

	// 调度的延迟定时器
	scheduleTimer = time.NewTimer(scheduleAfter)
	// 定时任务entities.Task
	for {
		select {
		case taskEvent = <- s.taskEventChan: // 监听任务变化事件
			//对内存中维护的任务列表做增删改查
			s.HandleTaskEvent(taskEvent)
		case <-  scheduleTimer.C: // 最近的任务到期了

		}
		// 调度一次任务
		scheduleAfter = s.TrySchedule()
		// 重置调度间隔
		scheduleTimer.Reset(scheduleAfter)
	}
}

// 处理任务事件
func (s *Scheduler) HandleTaskEvent(te *entities.TaskEvent)  {
	var (
		tsp *entities.TaskSchedulePlan
		isExist bool
		err error
	)
	switch te.EventType {
	case entities.SAVE: // 保存任务事情
		if tsp, err = entities.BuildTaskSchedulePlan(te.Task);err != nil {
			return
		}
		s.taskPlanTable[te.Task.Name] = tsp
	case entities.DELETE: // 删除任务事件
		if tsp, isExist = s.taskPlanTable[te.Task.Name]; isExist {
			delete(s.taskPlanTable, te.Task.Name)
		}
	}
}

// 推送任务变化事件
func (s *Scheduler) PublishTaskEvent(te  *entities.TaskEvent) {
	s.taskEventChan <- te
}

// 尝试执行任务
func (s *Scheduler) TryStartTask()  {

}


