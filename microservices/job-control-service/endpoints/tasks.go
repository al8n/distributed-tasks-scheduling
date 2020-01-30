package endpoints

import (
	"context"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/service"
	"github.com/go-kit/kit/endpoint"
)

func (ep *ImplEndpoints) MakeSaveOneTaskEndPoint(svc service.TasksManagerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return
	}
}

func (ep *ImplEndpoints) MakeDeleteOneTaskEndPoint(svc service.TasksManagerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return
	}
}

func (ep *ImplEndpoints) MakeGetOneTaskEndPoint(svc service.TasksManagerService) endpoint.Endpoint  {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return
	}
}

func (ep *ImplEndpoints) MakeGetAllTasksEndPoint(svc service.TasksManagerService) endpoint.Endpoint  {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return
	}
}

