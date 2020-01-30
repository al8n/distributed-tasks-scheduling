package requests

import "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/entities"

type SaveOneTaskRequest struct {
	Task entities.Task `json:"task"`
	Operator string `json:"operator"`
}

type DeleteOneTaskRequest struct {
	Name string `json:"name"`
}

type GetOneTaskRequest struct {
	Name string `json:"name"`
}

type GetAllTasksRequest struct {

}