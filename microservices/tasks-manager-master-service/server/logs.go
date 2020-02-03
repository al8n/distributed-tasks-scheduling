package server

import (
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/endpoints"
	grpcdecoder "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/grpc/decoder"
	grpcencoder "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/grpc/encoder"
	logspb "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/grpc/pb/logs"
	logsTransport "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/grpc/transport/logs_transport"
	httpdecoder "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/http/decoder"
	httpencoder "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/http/encoder"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/middleware/instrumentation"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/middleware/logging"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	httptransport "github.com/go-kit/kit/transport/http"
)

func (s *Server) LogsRoutes(f func(string) string)  {
	protol := HTTP_PROTOCOL
	var getLogsEndpoint endpoint.Endpoint
	{
		var method = "GetLogs"
		getLogsEndpoint = endpoints.SgtEndpoints.MakeGetAllLogsEndpoint(s.svc)
		getLogsEndpoint = instrumentation.SgtInstrumentation.InstrumentationMiddleware(protol, method)(getLogsEndpoint)
		getLogsEndpoint = logging.SgtLogging.LogMiddleware(protol, method)(getLogsEndpoint)
	}

	s.router.Methods("POST").Path(f("/logs")).Handler(
		httptransport.NewServer(
			getLogsEndpoint,
			httpdecoder.SgtHTTPDecoder.DecodeGetLogsRequest,
			httpencoder.SgtHTTPEncoder.EncodeGetLogsResponse,
		),
	)
}

func (s *Server) LogsGrpcRoutes () {
	protol := GRPC_PROTOCOL

	var getLogsEndpoint endpoint.Endpoint
	{
		var method = "GetLogs"
		getLogsEndpoint = endpoints.SgtEndpoints.MakeGetAllLogsEndpoint(s.svc)
		getLogsEndpoint = instrumentation.SgtInstrumentation.InstrumentationMiddleware(protol, method)(getLogsEndpoint)
		getLogsEndpoint = logging.SgtLogging.LogMiddleware(protol, method)(getLogsEndpoint)
	}

	s.gtp.Logs = &logsTransport.LogsTransport{
		GetLogsHandler: grpctransport.NewServer(
			getLogsEndpoint,
			grpcdecoder.SgtGRPCDecoder.DecodeLogsRequest,
			grpcencoder.SgtGRPCEncoder.EncodeGetLogsResponse,
		),
	}

	logspb.RegisterLogsServiceServer(s.gsrv, s.gtp.Logs)
}


