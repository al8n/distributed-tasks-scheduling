package server

import (
	myconfig "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/config"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/endpoints"
	grpcdecoder "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/grpc/decoder"
	grpcencoder "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/grpc/encoder"
	mygrpctransport "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/grpc/transport"
	httpdecoder "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/http/decoder"
	httpencoder "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/http/encoder"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/middleware/instrumentation"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/middleware/logging"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/repositories/distribution"
	mymongo "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/repositories/mongo"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/service"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Server struct {
	router *mux.Router
	svc    *service.ImplService
	srv	*http.Server
	gsrv   *grpc.Server
	gtp *mygrpctransport.ImplTasksTransport
}

var (
	SgtServer *Server
	once sync.Once
)

const (
	GRPC_PROTOCOL = "GRPC"
	HTTP_PROTOCOL = "HTTP"
)

func newServerSingleton() *Server  {
	once.Do(func() {
		r := mux.NewRouter()
		g := grpc.NewServer()


		SgtServer = &Server{
			router: r,
			svc:    service.InitService(),
			gsrv:   g,
			srv: &http.Server{
				Handler: r,
				Addr: ":" + strconv.Itoa(myconfig.ConfigSingleton.HTTPPort),
				WriteTimeout: time.Duration(myconfig.ConfigSingleton.HTTPWriteTimeout) * time.Millisecond ,
				ReadTimeout: time.Duration(myconfig.ConfigSingleton.HTTPReadTimeout) * time.Millisecond ,
			},
		}
	})
	return SgtServer
}

func v1(suffix string) string   {
	return "/api/v1" + suffix
}

func InitServer () (err error) {
	err = distribution.InitEtcd()
	if err != nil {
		log.Println(err)
	}

	err = mymongo.InitMongo()
	if err != nil {
		log.Println(err)
	}



	endpoints.InitEndpoints()
	instrumentation.InitInstrumentationMiddleware()
	logging.InitLoggingMiddleware()
	grpcdecoder.InitGRPCDecoder()
	grpcencoder.InitGRPCEncoder()
	httpdecoder.InitHTTPDecoder()
	httpencoder.InitHTTPEncoder()


	srv := newServerSingleton()

	gln, err := net.Listen("tcp", ":" + strconv.Itoa(myconfig.ConfigSingleton.GRPCPort))
	if err != nil {
		return
	}

	defer gln.Close()

	srv.router.Handle("/metrics", promhttp.Handler())
	srv.TasksGrpcRoutes()
	srv.TasksRoutes(v1)
	reflection.Register(SgtServer.gsrv)


	go func() {
		log.Fatal(srv.gsrv.Serve(gln))
	}()

	httpserver := &http.Server{
		Handler: srv.router,
		Addr: ":"+strconv.Itoa(myconfig.ConfigSingleton.HTTPPort),
		WriteTimeout: time.Duration(myconfig.ConfigSingleton.HTTPWriteTimeout),
		ReadTimeout: time.Duration(myconfig.ConfigSingleton.HTTPReadTimeout),
	}

	log.Fatal(httpserver.ListenAndServe())
	return
}