package endpoints

import (
	"context"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/requests"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/responses"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/services"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/utils/appErrors"
	"github.com/go-kit/kit/endpoint"
)

func (ep *ImplEndpoints) MakeSaveOneTaskEndPoint(svc services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(requests.SaveOneTaskRequest)
		if !ok {
			return nil, appErrors.EndpointAssertError
		}


		oldTask, err := svc.SaveOneTask(req.Task)
		if err != nil {
			return responses.Response{
				Errno: -1,
				Message: err.Error(),
			}, nil
		}
		return responses.Response{
			Errno:   0,
			Message: "",
			Data:    oldTask,
		}, nil
	}
}

func (ep *ImplEndpoints) MakeDeleteOneTaskEndPoint(svc services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(requests.DeleteOneTaskRequest)
		if !ok {
			return nil, appErrors.EndpointAssertError
		}

		oldTask, err := svc.DeleteOneTask(req.Name)
		if err != nil {
			return responses.Response{
				Errno: -1,
				Message: err.Error(),
			}, nil
		}
		return responses.Response{
			Errno:   0,
			Message: "",
			Data:    oldTask,
		}, nil
	}
}

func (ep *ImplEndpoints) MakeGetOneTaskEndPoint(svc services.Service) endpoint.Endpoint  {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(requests.GetOneTaskRequest)
		if !ok {
			return nil, appErrors.EndpointAssertError
		}

		task, err := svc.GetOneTask(req.Name)
		if err != nil {
			return responses.GetOneTaskResponse{
				Errno:   -1,
				Message: err.Error(),
				Task:    nil,
			}, nil
		}
		return responses.GetOneTaskResponse{
			Errno: 0,
			Message: "",
			Task: task,
		}, nil
	}
}

func (ep *ImplEndpoints) MakeGetAllTasksEndPoint(svc services.Service) endpoint.Endpoint  {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		_, ok := request.(requests.GetAllTasksRequest)
		if !ok {
			return nil, appErrors.EndpointAssertError
		}

		tasks, err := svc.GetAllTasks()
		if err != nil {
			return responses.GetAllTasksResponse{
				Errno:   -1,
				Message: err.Error(),
				Tasks:   nil,
			}, nil
		}
		return responses.GetAllTasksResponse{
			Errno:   0,
			Message: "",
			Tasks:   tasks,
		}, nil
	}
}

func (ep *ImplEndpoints) MakeKillOneTaskEndPoint(svc services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(requests.KillOneTaskRequest)
		if !ok {
			return nil, appErrors.EncoderAssertError
		}
		err = svc.KillOneTask(req.Name)
		if err != nil {
			return responses.Response{
				Errno:   -1,
				Message: err.Error(),
			}, nil
		}
		return responses.Response{
			Errno:   0,
			Message: "",
		}, nil
	}
}
