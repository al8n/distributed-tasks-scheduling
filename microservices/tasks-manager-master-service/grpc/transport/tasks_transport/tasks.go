package tasksTransport

import (
	"context"
	taskspb "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/grpc/pb/tasks"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type tasks interface {
	SaveOneTask(ctx context.Context, request *taskspb.SaveOneTaskRequest) (*taskspb.SaveOneTaskReply, error)
	DeleteOneTask(ctx context.Context, request *taskspb.DeleteOneTaskRequest) (*taskspb.DeleteOneTaskReply, error)
	KillOneTask(ctx context.Context, request *taskspb.DeleteOneTaskRequest) (*taskspb.KillOneTaskReply, error)
	GetOneTask(ctx context.Context, request *taskspb.GetOneTaskRequest) (*taskspb.GetOneTaskReply, error)
	GetAllTasks(ctx context.Context, request *taskspb.GetAllTasksRequest) (*taskspb.GetAllTasksReply, error)
}

type TasksTransport struct {
	SaveOneTaskHandler grpctransport.Handler
	DeleteOneTaskHandler grpctransport.Handler
	KillOneTaskHandler grpctransport.Handler
	GetOneTaskHandler grpctransport.Handler
	GetAllTasksHandler grpctransport.Handler
}

func (tp *TasksTransport) SaveOneTask(ctx context.Context, request *taskspb.SaveOneTaskRequest) (rpy *taskspb.SaveOneTaskReply, err error)  {
	_, rep, err := tp.SaveOneTaskHandler.ServeGRPC(ctx, request)
	if err != nil {
		return
	}

	return rep.(*taskspb.SaveOneTaskReply), nil
}

func (tp *TasksTransport) KillOneTask(ctx context.Context, request *taskspb.KillOneTaskRequest) (rpy *taskspb.KillOneTaskReply, err error)  {
	_, rep, err := tp.KillOneTaskHandler.ServeGRPC(ctx, request)
	if err != nil {
		return
	}

	return rep.(*taskspb.KillOneTaskReply), nil
}

func (tp *TasksTransport) DeleteOneTask(ctx context.Context, request *taskspb.DeleteOneTaskRequest) (rpy *taskspb.DeleteOneTaskReply, err error)  {
	_, rep, err := tp.DeleteOneTaskHandler.ServeGRPC(ctx, request)
	if err != nil {
		return
	}

	return rep.(*taskspb.DeleteOneTaskReply), nil
}

func (tp *TasksTransport) GetOneTask(ctx context.Context, request *taskspb.GetOneTaskRequest) (rpy *taskspb.GetOneTaskReply, err error)  {
	_, rep, err := tp.GetOneTaskHandler.ServeGRPC(ctx, request)
	if err != nil {
		return
	}

	return rep.(*taskspb.GetOneTaskReply), nil
}

func (tp *TasksTransport) GetAllTasks(ctx context.Context, request *taskspb.GetAllTasksRequest) (rpy *taskspb.
GetAllTasksReply, err error)  {
	_, rep, err := tp.GetAllTasksHandler.ServeGRPC(ctx, request)
	if err != nil {
		return
	}

	return rep.(*taskspb.GetAllTasksReply), nil
}


