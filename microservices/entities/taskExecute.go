package entities

import "time"

// 任务执行状态
type TaskExecuteInfo struct {
	Task *Task // 任务信息
	PlanTime time.Time // 理论上的调度时间
	RealTime time.Time // 实际的调度时间
}

// 构造执行状态信息
func BuildTaskExecuteInfo(tsp *TaskSchedulePlan)  (tei *TaskExecuteInfo) {
	tei = &TaskExecuteInfo{
		Task: tsp.Task,
		PlanTime: tsp.NextTime, //计算调度时间
		RealTime: time.Now(), //真实调度时间
	}
	return
}

// 任务执行结果
type TaskExecuteResult struct {
	ExecuteInfo *TaskExecuteInfo
	Output []byte //脚本输出
	Err error
	StartTime time.Time
	EndTime time.Time
}