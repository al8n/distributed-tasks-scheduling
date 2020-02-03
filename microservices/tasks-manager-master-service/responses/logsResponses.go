package responses

import "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/entities"

type GetLogsResponse struct {
	Errno int32  `json:"error_number"`
	Message string `json:"msg"`
	Logs []*entities.TaskLog `json:"logs"`
}
