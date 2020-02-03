package logsTransport

import (
	"context"
	logspb "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/grpc/pb/logs"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type logs interface {
	GetLogs(ctx context.Context, request *logspb.GetLogsRequest) (*logspb.GetLogsReply, error)
}

type LogsTransport struct {
	GetLogsHandler grpctransport.Handler
}

func (tp *LogsTransport) GetLogs(ctx context.Context, request *logspb.GetLogsRequest) (rpy *logspb.GetLogsReply,
	err error)  {
	_, rep, err := tp.GetLogsHandler.ServeGRPC(ctx, request)
	if err != nil {
		return
	}

	return rep.(*logspb.GetLogsReply), nil
}