package entities

import (
	"encoding/json"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/conf"
	"github.com/gorhill/cronexpr"
	"strings"
	"time"
)

type Task struct {
	Name string `json:"name"`
	Command string `json:"command"`
	Expression string  `json:"expression"`
}

// 反序列化task
func UnMarshallTask(value  []byte) (rst *Task, err error) {
	var (
		task *Task
	)

	task = &Task{}
	if err = json.Unmarshal(value, task); err != nil {
		return
	}
	rst = task
	return
}

// 从etcd的key中提取任务名
func ExtractTaskName(key string) string  {
	return strings.TrimPrefix(key, conf.TASK_SAVE_DIR)
}

type TaskSchedulePlan struct {
	Task *Task
	Expression *cronexpr.Expression // 解析好的cronexpr表达式
	NextTime time.Time // 下次调度时间
}

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