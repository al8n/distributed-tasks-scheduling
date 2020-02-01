package entities

import (
	"github.com/gorhill/cronexpr"
	"encoding/json"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/conf"
	"strings"
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
	Expression *cronexpr.Expression
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
	task *Task
}

// 任务变化事件有2种：1）更新任务 2）删除任务
func BuildTaskEvent(et EventType, task *Task) (te *TaskEvent) {
	return &TaskEvent{
		EventType: et,
		task: task,
	}
}