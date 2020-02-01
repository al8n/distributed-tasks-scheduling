package entities

import (
	"github.com/gorhill/cronexpr"
	"time"
)

type EventType int

const (
	SAVE EventType = iota
	DELETE
	KILL
)

// 变化事件
type TaskEvent struct {
	EventType EventType // Save, Delete
	Task *Task
}

// 任务执行计划
type TaskSchedulePlan struct {
	Task *Task
	Expression *cronexpr.Expression // 解析好的cronexpr表达式
	NextTime time.Time // 下次调度时间
}


// 任务变化事件有2种：1）更新任务 2）删除任务
func BuildTaskEvent(et EventType, task *Task) (te *TaskEvent) {
	return &TaskEvent{
		EventType: et,
		Task: task,
	}
}

// 构造任务执行计划
func BuildTaskSchedulePlan(task *Task) (tsp *TaskSchedulePlan, err error) {
	var (
		expr *cronexpr.Expression
	)

	// 解析task的cron表达式
	if expr, err = cronexpr.Parse(task.Expression); err != nil {
		return
	}

	// 生成任务调度计划对象
	tsp = &TaskSchedulePlan{
		Task: task,
		Expression: expr,
		NextTime: expr.Next(time.Now()),
	}
	return
}