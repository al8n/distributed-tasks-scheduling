package endpoints

import (
	"context"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/requests"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/responses"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/service"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/utils/appErrors"
	"github.com/go-kit/kit/endpoint"
)

func (ep *ImplEndpoints) MakeSaveOneTaskEndPoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(requests.SaveOneTaskRequest)
		if !ok {
			return nil, appErrors.AssertError
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

func (ep *ImplEndpoints) MakeDeleteOneTaskEndPoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(requests.DeleteOneTaskRequest)
		if !ok {
			return nil, appErrors.AssertError
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

func (ep *ImplEndpoints) MakeGetOneTaskEndPoint(svc service.Service) endpoint.Endpoint  {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(requests.GetOneTaskRequest)
		if !ok {
			return nil, appErrors.AssertError
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

func (ep *ImplEndpoints) MakeGetAllTasksEndPoint(svc service.Service) endpoint.Endpoint  {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		_, ok := request.(requests.GetAllTasksRequest)
		if !ok {
			return nil, appErrors.AssertError
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

