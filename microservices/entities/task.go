package entities

import (
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









