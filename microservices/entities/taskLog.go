package entities

type TaskLog struct {
	TaskName string `bson:"task_name" json:"task_name"` // 任务名字
	Command string `bson:"command" json:"command"` // 脚本命令
	Error string `bson:"error" json:"error"` // 错误原因
	Output string `bson:"output" json:"output"` // 脚本输出
	PlanTime int64 `bson:"plan_time" json:"plan_time"` // 计划开始时间
	ScheduleTime int64 `bson:"schedule_time" json:"schedule_time"` // 实际调度时间
	StartTime int64 `bson:"start_time" json:"start_time"` // 任务执行开始时间
	EndTime int64 `bson:"end_time" json:"end_time"` // 任务执行结束时间
}
