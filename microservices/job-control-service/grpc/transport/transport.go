package transport

import (
	"context"
	taskspb "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/grpc/pb/tasks"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type Transports interface {
	SaveOneTask(ctx context.Context, request *taskspb.SaveOneTaskRequest) (*taskspb.OneTaskReply, error)
	DeleteOneTask(ctx context.Context, request *taskspb.DeleteOneTaskRequest) (*taskspb.OneTaskReply, error)
	GetOneTask(ctx context.Context, request *taskspb.GetOneTaskRequest) (*taskspb.GetOneTaskReply, error)
	GetAllTasks(ctx context.Context, request *taskspb.GetAllTasksRequest) (*taskspb.GetAllTasksReply, error)
}

type ImplTasksTransport struct {
	SaveOneTaskHandler grpctransport.Handler
	DeleteOneTaskHandler grpctransport.Handler
	GetOneTaskHandler grpctransport.Handler
	GetAllTasksHandler grpctransport.Handler
}


func (tp *ImplTasksTransport) SaveOneTask(ctx context.Context, request *taskspb.SaveOneTaskRequest) (rpy *taskspb.OneTaskReply, err error)  {
	_, rep, err := tp.SaveOneTaskHandler.ServeGRPC(ctx, request)
	if err != nil {
		return
	}

	return rep.(*taskspb.OneTaskReply), nil
}

func (tp *ImplTasksTransport) DeleteOneTask(ctx context.Context, request *taskspb.DeleteOneTaskRequest) (rpy *taskspb.
OneTaskReply, err error)  {
	_, rep, err := tp.DeleteOneTaskHandler.ServeGRPC(ctx, request)
	if err != nil {
		return
	}

	return rep.(*taskspb.OneTaskReply), nil
}

func (tp *ImplTasksTransport) GetOneTask(ctx context.Context, request *taskspb.GetOneTaskRequest) (rpy *taskspb.GetOneTaskReply, err error)  {
	_, rep, err := tp.GetOneTaskHandler.ServeGRPC(ctx, request)
	if err != nil {
		return
	}

	return rep.(*taskspb.GetOneTaskReply), nil
}

func (tp *ImplTasksTransport) GetAllTasks(ctx context.Context, request *taskspb.GetAllTasksRequest) (rpy *taskspb.
GetAllTasksReply, err error)  {
	_, rep, err := tp.GetAllTasksHandler.ServeGRPC(ctx, request)
	if err != nil {
		return
	}

	return rep.(*taskspb.GetAllTasksReply), nil
}

