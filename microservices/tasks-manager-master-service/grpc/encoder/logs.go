package encoder

import (
	"context"
	logspb "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/grpc/pb/logs"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/responses"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/utils"
)

func (gec *ImplEncoder) EncodeGetLogsResponse(_ context.Context, response interface{}) (interface{}, error)  {
	var (
		logs []*logspb.Logs
	)
	pbrpl := &logspb.GetLogsReply{}
	res, ok := response.(responses.GetLogsResponse)
	if !ok {
		pbrpl.ErrorNumber = -1
		pbrpl.Msg = utils.EncoderAssertError.Error()
		return pbrpl, nil
	}

	for _, log := range res.Logs {
		logs = append(logs, &logspb.Logs{
			TaskName:             log.TaskName,
			Command:              log.Command,
			Error:                log.Error,
			Output:               log.Output,
			PlanTime:             log.PlanTime,
			ScheduleTime:         log.ScheduleTime,
			StartTime:            log.StartTime,
			EndTime:              log.EndTime,
		})
	}
	pbrpl.Logs = logs
	pbrpl.Msg = ""
	pbrpl.ErrorNumber = 0
	return pbrpl, nil
}
