package endpoints

import (
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/services"
	"github.com/go-kit/kit/endpoint"
	"sync"
)

type Endpoints interface {
	MakeSaveOneTaskEndPoint(svc services.Service) endpoint.Endpoint
	MakeDeleteOneTaskEndPoint(svc services.Service) endpoint.Endpoint
	MakeKillOneTaskEndPoint(svc services.Service) endpoint.Endpoint
	MakeGetOneTaskEndPoint(svc services.Service) endpoint.Endpoint
	MakeGetAllTasksEndPoint(svc services.Service) endpoint.Endpoint
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

