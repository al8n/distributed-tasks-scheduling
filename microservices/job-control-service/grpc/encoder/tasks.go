package encoder

import (
	"context"
	taskspb "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/grpc/pb/tasks"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/responses"
)

func (gec *ImplEncoder) EncodeSaveOneTaskResponse(_ context.Context, response interface{}) (i interface{}, err error) {
	res, ok := response.(responses.Response)
	if !ok {
		return
	}

	return &taskspb.OneTaskReply{
		ErrorNumber: 0,
		Msg: res.Message,
	}, nil
}

func (gec *ImplEncoder) EncodeDeleteOneTaskResponse(_ context.Context, response interface{}) (i interface{}, err error) {
	res, ok := response.(responses.Response)
	if !ok {
		return
	}

	return &taskspb.OneTaskReply{
		ErrorNumber: 0,
		Msg: res.Message,
	}, nil
}

func (gec *ImplEncoder) EncodeGetOneTaskResponse(_ context.Context, response interface{}) (i interface{}, err error) {
	res, ok := response.(responses.GetOneTaskResponse)
	if !ok {
		return
	}

	return &taskspb.GetOneTaskReply{
		ErrorNumber: 0,
		Msg: res.Message,
	}, nil
}

func (gec *ImplEncoder) EncodeGetAllTasksResponse(_ context.Context, response interface{}) (i interface{}, err error) {
	var tasks []*taskspb.Task

	res, ok := response.(responses.GetAllTasksResponse)
	if !ok {
		return
	}

	for _, task := range res.Tasks {
		tasks = append(tasks, &taskspb.Task{
			Name:                 task.Name,
			Command:              task.Command,
			Expression:           task.Expression,
		})
	}
	
	return &taskspb.GetAllTasksReply{
		Tasks: tasks,
		ErrorNumber: 0,
		Msg: res.Message,
	}, nil
}
