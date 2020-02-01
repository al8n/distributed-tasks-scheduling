package endpoints

import (
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/service"
	"github.com/go-kit/kit/endpoint"
	"sync"
)

type Endpoints interface {
	MakeSaveOneTaskEndPoint(svc service.Service) endpoint.Endpoint
	MakeDeleteOneTaskEndPoint(svc service.Service) endpoint.Endpoint
	MakeKillOneTaskEndPoint(svc service.Service) endpoint.Endpoint
	MakeGetOneTaskEndPoint(svc service.Service) endpoint.Endpoint
	MakeGetAllTasksEndPoint(svc service.Service) endpoint.Endpoint
}

type ImplEndpoints struct {
}


var  (
	SgtEndpoints *ImplEndpoints
	once sync.Once
)

func InitEndpoints() *ImplEndpoints  {
	once.Do(func() {
		SgtEndpoints = &ImplEndpoints{}
	})
	return  SgtEndpoints
}

