package requests

import "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/entities"

type SaveOneTaskRequest struct {
	Task     entities.Task `json:"task"`
	Operator string        `json:"operator"`
}

type DeleteOneTaskRequest struct {
	Name string `json:"name"`
}

type KillOneTaskRequest struct {
	Name string `json:"name"`
}

type GetOneTaskRequest struct {
	Name string `json:"name"`
}

type GetAllTasksRequest struct {}