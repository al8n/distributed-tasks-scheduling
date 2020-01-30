package endpoints

import (
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/service"
	"github.com/go-kit/kit/endpoint"
	"sync"
)

type endpoints interface {
	MakeSaveOneTaskEndPoint(svc service.TasksManagerService) endpoint.Endpoint
	MakeDeleteOneTaskEndPoint(svc service.TasksManagerService) endpoint.Endpoint
	MakeGetOneTaskEndPoint(svc service.TasksManagerService) endpoint.Endpoint
	MakeGetAllTasksEndPoint(svc service.TasksManagerService) endpoint.Endpoint
}

type ImplEndpoints struct {}

var (
	EndpointsSingleton *ImplEndpoints
	once sync.Once
)

func NewEndpoints() *ImplEndpoints  {
	once.Do(func() {
		EndpointsSingleton = &ImplEndpoints{}
	})

	return EndpointsSingleton
}

