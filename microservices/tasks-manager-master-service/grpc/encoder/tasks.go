package encoder

import (
	"context"
	taskspb "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/grpc/pb/tasks"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/responses"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/utils/appErrors"
)

func (gec *ImplEncoder) EncodeSaveOneTaskResponse(_ context.Context, response interface{}) (i interface{}, err error) {
	pbrpl := &taskspb.SaveOneTaskReply{}
	res, ok := response.(responses.Response)
	if !ok {
		pbrpl.ErrorNumber = -1
		pbrpl.Msg = appErrors.EncoderAssertError.Error()
		return pbrpl, nil
	}
	pbrpl.Msg = res.Message
	pbrpl.ErrorNumber = 0
	return pbrpl, nil
}

func (gec *ImplEncoder) EncodeDeleteOneTaskResponse(_ context.Context, response interface{}) (i interface{}, err error) {
	pbrpl := &taskspb.DeleteOneTaskReply{}
	res, ok := response.(responses.Response)
	if !ok {
		pbrpl.ErrorNumber = -1
		pbrpl.Msg = appErrors.EncoderAssertError.Error()
		return pbrpl, nil
	}

	pbrpl.Msg = res.Message
	pbrpl.ErrorNumber = 0

	return pbrpl, nil
}

func (gec *ImplEncoder) EncodeGetOneTaskResponse(_ context.Context, response interface{}) (i interface{}, err error) {
	pbrpl := &taskspb.GetOneTaskReply{}

	res, ok := response.(responses.GetOneTaskResponse)

	if !ok {
		pbrpl.ErrorNumber = -1
		pbrpl.Msg = appErrors.EncoderAssertError.Error()
		return pbrpl, nil
	}

	pbrpl.Msg = res.Message
	pbrpl.ErrorNumber = 0
	pbrpl.Task = &taskspb.Task{
		Name:                 res.Task.Name,
		Command:              res.Task.Command,
		Expression:           res.Task.Expression,
	}


	return pbrpl, nil
}

func (gec *ImplEncoder) EncodeGetAllTasksResponse(_ context.Context, response interface{}) (i interface{}, err error) {
	var (
		tasks []*taskspb.Task
		pbrpl *taskspb.GetAllTasksReply
	)

	res, ok := response.(responses.GetAllTasksResponse)
	if !ok {
		pbrpl = &taskspb.GetAllTasksReply{
			ErrorNumber:          -1,
			Msg:                  appErrors.EncoderAssertError.Error(),
		}
		return
	}

	for _, task := range res.Tasks {
		tasks = append(tasks, &taskspb.Task{
			Name:                 task.Name,
			Command:              task.Command,
			Expression:           task.Expression,
		})
	}

	pbrpl = &taskspb.GetAllTasksReply{
		Tasks: tasks,
		ErrorNumber: 0,
		Msg: res.Message,
	}

	return pbrpl, nil
}

func (gec *ImplEncoder) EncodeKillOneTaskResponse(_ context.Context, response interface{}) (interface{}, error) {
	pbrpl := &taskspb.KillOneTaskReply{}
	res, ok := response.(responses.Response)
	if !ok {
		pbrpl.ErrorNumber = -1
		pbrpl.Msg = appErrors.EncoderAssertError.Error()
		return pbrpl, nil
	}
	pbrpl.Msg = res.Message
	pbrpl.ErrorNumber = 0
	return pbrpl, nil
}