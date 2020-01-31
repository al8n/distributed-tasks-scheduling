package server

import (
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/endpoints"
	grpcdecoder "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/grpc/decoder"
	grpcencoder "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/grpc/encoder"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/grpc/transport"
	httpdecoder "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/http/decoder"
	httpencoder "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/http/encoder"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/middleware/logging"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	httptransport "github.com/go-kit/kit/transport/http"
)

func (s *Server) TasksRoutes(f func(string) string)  {

	s.router.Methods("POST").Path(f("/task/save/one")).Handler(
		httptransport.NewServer(
			logging.SgtLogging.TaskLoggingMiddleware(HTTP_PROTOCOL, "SaveOneTask")(endpoints.SgtEndpoints.MakeSaveOneTaskEndPoint(s.svc)),
			httpdecoder.SgtHTTPDecoder.DecodeSaveOneTaskRequest,
			httpencoder.SgtHTTPEncoder.EncodeSaveOneTaskResponse,
		),
	)

	s.router.Methods("POST").Path(f("/task/delete/one")).Handler(
		httptransport.NewServer(
			logging.SgtLogging.TaskLoggingMiddleware(HTTP_PROTOCOL, "DeleteOneTask")(endpoints.SgtEndpoints.MakeDeleteOneTaskEndPoint(s.svc)),
			httpdecoder.SgtHTTPDecoder.DecodeDeleteOneTaskRequest,
			httpencoder.SgtHTTPEncoder.EncodeDeleteOneTaskResponse,
		),
	)

	s.router.Methods("GET").Path(f("/task/get/one/{name}")).Handler(
		httptransport.NewServer(
			logging.SgtLogging.TaskLoggingMiddleware(HTTP_PROTOCOL, "GetOneTask")(endpoints.SgtEndpoints.MakeGetOneTaskEndPoint(s.svc)),
			httpdecoder.SgtHTTPDecoder.DecodeGetOneTaskRequest,
			httpencoder.SgtHTTPEncoder.EncodeGetOneTaskResponse,
		),
	)

	s.router.Methods("GET").Path(f("/task/get/all")).Handler(
		httptransport.NewServer(
			logging.SgtLogging.TaskLoggingMiddleware(HTTP_PROTOCOL, "GetAllTasks")(endpoints.SgtEndpoints.MakeGetAllTasksEndPoint(s.svc)),
			httpdecoder.SgtHTTPDecoder.DecodeGetAllTasksRequest,
			httpencoder.SgtHTTPEncoder.EncodeGetAllTasksResponse,
		),
	)

}


func (s *Server) TasksGrpcRoutes() {

	s.svc.Transports = &transport.ImplTasksTransport{
		SaveOneTaskHandler:   grpctransport.NewServer(
			logging.SgtLogging.TaskLoggingMiddleware(GRPC_PROTOCOL, "SaveOneTask")(endpoints.SgtEndpoints.MakeGetOneTaskEndPoint(s.
				svc)),
			grpcdecoder.SgtGRPCDecoder.DecodeSaveOneTaskRequest,
			grpcencoder.SgtGRPCEncoder.EncodeSaveOneTaskResponse,
		),
		DeleteOneTaskHandler: grpctransport.NewServer(
			logging.SgtLogging.TaskLoggingMiddleware(GRPC_PROTOCOL, "DeleteOneTask")(endpoints.SgtEndpoints.MakeDeleteOneTaskEndPoint(s.svc)),
			grpcdecoder.SgtGRPCDecoder.DecodeDeleteOneTaskRequest,
			grpcencoder.SgtGRPCEncoder.EncodeDeleteOneTaskResponse,
		),
		GetOneTaskHandler:    grpctransport.NewServer(
			logging.SgtLogging.TaskLoggingMiddleware(GRPC_PROTOCOL, "GetOneTask")(endpoints.SgtEndpoints.MakeGetOneTaskEndPoint(s.svc)),
			grpcdecoder.SgtGRPCDecoder.DecodeGetOneTaskRequest,
			grpcencoder.SgtGRPCEncoder.EncodeGetOneTaskResponse,
		),
		GetAllTasksHandler:   grpctransport.NewServer(
			logging.SgtLogging.TaskLoggingMiddleware(GRPC_PROTOCOL, "GetAllTasks")(endpoints.SgtEndpoints.MakeGetAllTasksEndPoint(s.svc)),
			grpcdecoder.SgtGRPCDecoder.DecodeGetAllTasksRequest,
			grpcencoder.SgtGRPCEncoder.EncodeGetAllTasksResponse,
		),
	}
}