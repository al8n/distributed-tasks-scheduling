package server

import (
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/endpoints"
	grpcdecoder "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/grpc/decoder"
	grpcencoder "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/grpc/encoder"
	taskspb "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/grpc/pb/tasks"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/grpc/transport"
	httpdecoder "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/http/decoder"
	httpencoder "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/http/encoder"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/middleware/instrumentation"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/middleware/logging"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	httptransport "github.com/go-kit/kit/transport/http"
)

func (s *Server) TasksRoutes(f func(string) string)  {

	protol := HTTP_PROTOCOL
	var saveOneTaskEndpoint endpoint.Endpoint
	{
		var method = "SaveOneTask"
		saveOneTaskEndpoint = endpoints.SgtEndpoints.MakeSaveOneTaskEndPoint(s.svc)
		saveOneTaskEndpoint = instrumentation.SgtInstrumentation.TaskInstrumentationMiddleware(protol,method)(saveOneTaskEndpoint)
		saveOneTaskEndpoint = logging.SgtLogging.TaskLoggingMiddleware(protol, method)(saveOneTaskEndpoint)
	}

	var deleteOneTaskEndpoint endpoint.Endpoint
	{
		var method = "DeleteOneTask"
		deleteOneTaskEndpoint = endpoints.SgtEndpoints.MakeDeleteOneTaskEndPoint(s.svc)
		deleteOneTaskEndpoint = instrumentation.SgtInstrumentation.TaskInstrumentationMiddleware(protol, method)(deleteOneTaskEndpoint)
		deleteOneTaskEndpoint = logging.SgtLogging.TaskLoggingMiddleware(protol, method)(deleteOneTaskEndpoint)
	}

	var getOneTaskEndpoint endpoint.Endpoint
	{
		var method = "GetOneTask"
		getOneTaskEndpoint = endpoints.SgtEndpoints.MakeGetOneTaskEndPoint(s.svc)
		getOneTaskEndpoint = instrumentation.SgtInstrumentation.TaskInstrumentationMiddleware(protol, method)(getOneTaskEndpoint)
		getOneTaskEndpoint = logging.SgtLogging.TaskLoggingMiddleware(protol, method)(getOneTaskEndpoint)
	}

	var getAllTasksEndpoint endpoint.Endpoint
	{
		var method = "GetAllTasks"
		getAllTasksEndpoint = endpoints.SgtEndpoints.MakeGetAllTasksEndPoint(s.svc)
		getAllTasksEndpoint = instrumentation.SgtInstrumentation.TaskInstrumentationMiddleware(protol, method)(getAllTasksEndpoint)
		getAllTasksEndpoint = logging.SgtLogging.TaskLoggingMiddleware(protol, method)(getAllTasksEndpoint)
	}

	var killOneTaskEndpoint endpoint.Endpoint
	{
		var method = "KillOneTask"
		killOneTaskEndpoint = endpoints.SgtEndpoints.MakeKillOneTaskEndPoint(s.svc)
		killOneTaskEndpoint = instrumentation.SgtInstrumentation.TaskInstrumentationMiddleware(protol, method)(killOneTaskEndpoint)
		killOneTaskEndpoint = logging.SgtLogging.TaskLoggingMiddleware(protol, method)(killOneTaskEndpoint)
	}

	s.router.Methods("POST").Path(f("/task/save/one")).Handler(
		httptransport.NewServer(
			saveOneTaskEndpoint,
			httpdecoder.SgtHTTPDecoder.DecodeSaveOneTaskRequest,
			httpencoder.SgtHTTPEncoder.EncodeSaveOneTaskResponse,
		),
	)

	s.router.Methods("POST").Path(f("/task/delete/one")).Handler(
		httptransport.NewServer(
			deleteOneTaskEndpoint,
			httpdecoder.SgtHTTPDecoder.DecodeDeleteOneTaskRequest,
			httpencoder.SgtHTTPEncoder.EncodeDeleteOneTaskResponse,
		),
	)

	s.router.Methods("POST").Path(f("/task/kill/one")).Handler(
		httptransport.NewServer(
			killOneTaskEndpoint,
			httpdecoder.SgtHTTPDecoder.DecodeKillOneTaskRequest,
			httpencoder.SgtHTTPEncoder.EncodeKillOneTaskResponse,
		),
	)

	s.router.Methods("POST").Path(f("/task/get/one")).Handler(
		httptransport.NewServer(
			getOneTaskEndpoint,
			httpdecoder.SgtHTTPDecoder.DecodeGetOneTaskRequest,
			httpencoder.SgtHTTPEncoder.EncodeGetOneTaskResponse,
		),
	)

	s.router.Methods("GET").Path(f("/task/get/all")).Handler(
		httptransport.NewServer(
			getAllTasksEndpoint,
			httpdecoder.SgtHTTPDecoder.DecodeGetAllTasksRequest,
			httpencoder.SgtHTTPEncoder.EncodeGetAllTasksResponse,
		),
	)

}

func (s *Server) TasksGrpcRoutes() {
	protol := GRPC_PROTOCOL
	var saveOneTaskEndpoint endpoint.Endpoint
	{
		var method = "SaveOneTask"
		saveOneTaskEndpoint = endpoints.SgtEndpoints.MakeSaveOneTaskEndPoint(s.svc)
		saveOneTaskEndpoint = instrumentation.SgtInstrumentation.TaskInstrumentationMiddleware(protol,method)(saveOneTaskEndpoint)
		saveOneTaskEndpoint = logging.SgtLogging.TaskLoggingMiddleware(protol, method)(saveOneTaskEndpoint)
	}

	var deleteOneTaskEndpoint endpoint.Endpoint
	{
		var method = "DeleteOneTask"
		deleteOneTaskEndpoint = endpoints.SgtEndpoints.MakeDeleteOneTaskEndPoint(s.svc)
		deleteOneTaskEndpoint = instrumentation.SgtInstrumentation.TaskInstrumentationMiddleware(protol, method)(deleteOneTaskEndpoint)
		deleteOneTaskEndpoint = logging.SgtLogging.TaskLoggingMiddleware(protol, method)(deleteOneTaskEndpoint)
	}

	var getOneTaskEndpoint endpoint.Endpoint
	{
		var method = "GetOneTask"
		getOneTaskEndpoint = endpoints.SgtEndpoints.MakeGetOneTaskEndPoint(s.svc)
		getOneTaskEndpoint = instrumentation.SgtInstrumentation.TaskInstrumentationMiddleware(protol, method)(getOneTaskEndpoint)
		getOneTaskEndpoint = logging.SgtLogging.TaskLoggingMiddleware(protol, method)(getOneTaskEndpoint)
	}

	var getAllTasksEndpoint endpoint.Endpoint
	{
		var method = "GetAllTasks"
		getAllTasksEndpoint = endpoints.SgtEndpoints.MakeGetAllTasksEndPoint(s.svc)
		getAllTasksEndpoint = instrumentation.SgtInstrumentation.TaskInstrumentationMiddleware(protol, method)(getAllTasksEndpoint)
		getAllTasksEndpoint = logging.SgtLogging.TaskLoggingMiddleware(protol, method)(getAllTasksEndpoint)
	}

	var killOneTaskEndpoint endpoint.Endpoint
	{
		var method = "KillOneTask"
		killOneTaskEndpoint = endpoints.SgtEndpoints.MakeKillOneTaskEndPoint(s.svc)
		killOneTaskEndpoint = instrumentation.SgtInstrumentation.TaskInstrumentationMiddleware(protol, method)(killOneTaskEndpoint)
		killOneTaskEndpoint = logging.SgtLogging.TaskLoggingMiddleware(protol, method)(killOneTaskEndpoint)
	}
	s.gtp = &transport.ImplTasksTransport{
		SaveOneTaskHandler:   grpctransport.NewServer(
			saveOneTaskEndpoint,
			grpcdecoder.SgtGRPCDecoder.DecodeSaveOneTaskRequest,
			grpcencoder.SgtGRPCEncoder.EncodeSaveOneTaskResponse,
		),
		DeleteOneTaskHandler: grpctransport.NewServer(
			deleteOneTaskEndpoint,
			grpcdecoder.SgtGRPCDecoder.DecodeDeleteOneTaskRequest,
			grpcencoder.SgtGRPCEncoder.EncodeDeleteOneTaskResponse,
		),
		KillOneTaskHandler: grpctransport.NewServer(
			killOneTaskEndpoint,
			grpcdecoder.SgtGRPCDecoder.DecodeKillOneTaskRequest,
			grpcencoder.SgtGRPCEncoder.EncodeKillOneTaskResponse,
		),
		GetOneTaskHandler:    grpctransport.NewServer(
			getOneTaskEndpoint,
			grpcdecoder.SgtGRPCDecoder.DecodeGetOneTaskRequest,
			grpcencoder.SgtGRPCEncoder.EncodeGetOneTaskResponse,
		),
		GetAllTasksHandler:   grpctransport.NewServer(
			getAllTasksEndpoint,
			grpcdecoder.SgtGRPCDecoder.DecodeGetAllTasksRequest,
			grpcencoder.SgtGRPCEncoder.EncodeGetAllTasksResponse,
		),
	}

	taskspb.RegisterTasksServer(s.gsrv, s.gtp)

}