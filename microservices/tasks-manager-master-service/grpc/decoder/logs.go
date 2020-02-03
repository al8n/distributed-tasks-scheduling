package decoder

import (
	"context"
	logspb "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/grpc/pb/logs"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/requests"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/utils"
)

func(gdc *ImplDecoder) DecodeLogsRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*logspb.GetLogsRequest)
	if !ok {
		return nil, utils.DecoderAssertError
	}

	return requests.GetLogsRequest{
		Start:         req.Start,
		After:         req.After,
		Page:    req.PageNumber,
		Limit: req.ResultPerPage,
		Field:         req.Field,
	}, nil
}
