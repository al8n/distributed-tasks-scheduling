package endpoints

import (
	"context"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/requests"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/responses"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/services"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/utils"
	"github.com/go-kit/kit/endpoint"
)

func (ep *ImplEndpoints) MakeGetAllLogsEndpoint(svc services.Service) endpoint.Endpoint  {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(requests.GetLogsRequest)
		if !ok {
			return nil, utils.EndpointAssertError
		}
		logs, err := svc.GetLogs(req)
		if err != nil {
			return responses.GetLogsResponse{
				Errno:   -1,
				Message: err.Error(),
				Logs:    nil,
			}, nil
		}

		return responses.GetLogsResponse{
			Errno:   0,
			Message: "",
			Logs:    logs,
		}, nil
	}
}
