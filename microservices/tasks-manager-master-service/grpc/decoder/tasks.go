package decoder

import (
	"context"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/entities"
	taskspb "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/grpc/pb/tasks"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/requests"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/utils/appErrors"
	"log"
)

func (gdc *ImplDecoder) DecodeSaveOneTaskRequest(ctx context.Context, grpcReq interface{}) (i interface{},
err error) {
	req, ok := grpcReq.(*taskspb.SaveOneTaskRequest)
	if !ok {
		return nil, appErrors.DecoderAssertError
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
		return nil, appErrors.DecoderAssertError
	}

	return requests.DeleteOneTaskRequest{Name: req.TaskKey}, nil
}

func (gdc *ImplDecoder) DecodeGetAllTasksRequest(ctx context.Context, grpcReq interface{}) (i interface{}, err error) {
	req, ok := grpcReq.(*taskspb.GetAllTasksRequest)
	if !ok {
		return nil, appErrors.DecoderAssertError
	}
	log.Println(req.Token)
	return requests.GetAllTasksRequest{}, nil
}

func (gdc *ImplDecoder) DecodeGetOneTaskRequest(ctx context.Context, grpcReq interface{}) (i interface{}, err error) {
	req, ok := grpcReq.(*taskspb.GetOneTaskRequest)
	if !ok {
		return nil, appErrors.DecoderAssertError
	}
	return  requests.GetOneTaskRequest{Name: req.TaskKey}, nil
}

func (gdc *ImplDecoder) DecodeKillOneTaskRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*taskspb.KillOneTaskRequest)
	if !ok {
		return nil, appErrors.DecoderAssertError
	}
	return requests.KillOneTaskRequest{
		Name: req.TaskKey,
	}, nil
}