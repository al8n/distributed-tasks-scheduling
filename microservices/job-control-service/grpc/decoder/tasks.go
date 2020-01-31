package decoder

import (
	"context"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/entities"
	taskspb "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/grpc/pb/tasks"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/requests"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/utils/appErrors"
	"log"
)

func (gdc *ImplDecoder) DecodeSaveOneTaskRequest(ctx context.Context, grpcReq interface{}) (i interface{},
err error) {
	req, ok := grpcReq.(*taskspb.SaveOneTaskRequest)
	if !ok {
		return nil, appErrors.AssertError
	}

	return requests.SaveOneTaskRequest{
		Task: entities.Task{
			Name:       req.Task.Name,
			Command:    req.Task.Command,
			Expression: req.Task.Expression,
		},
	}, nil
}


func (gdc *ImplDecoder) DecodeDeleteOneTaskRequest(ctx context.Context, grpcReq interface{}) (i interface{},
err error) {
	req, ok := grpcReq.(*taskspb.DeleteOneTaskRequest)
	if !ok {
		return nil, appErrors.AssertError
	}

	return requests.DeleteOneTaskRequest{Name: req.TaskKey}, nil
}

func (gdc *ImplDecoder) DecodeGetAllTasksRequest(ctx context.Context, grpcReq interface{}) (i interface{}, err error) {
	req, ok := grpcReq.(*taskspb.GetAllTasksRequest)
	if !ok {
		return nil, appErrors.AssertError
	}
	log.Println(req.Token)
	return requests.GetAllTasksRequest{}, nil
}

func (gdc *ImplDecoder) DecodeGetOneTaskRequest(ctx context.Context, grpcReq interface{}) (i interface{}, err error) {
	req, ok := grpcReq.(*taskspb.GetOneTaskRequest)
	if !ok {
		return nil, appErrors.AssertError
	}
	return  requests.GetOneTaskRequest{Name: req.TaskKey}, nil
}