package server

import (
	myconfig "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/config"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/endpoints"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/repositories/distribution"
	mymongo "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/repositories/mongo"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/service"
	"github.com/gorilla/mux"
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
	svc    *service.TasksManagerImpl
	srv	*http.Server
	gsrv   *grpc.Server
	ep *endpoints.ImplEndpoints
}

var (
	ServerSingleton *Server
	once sync.Once
)

func newServerSingleton() *Server  {
	once.Do(func() {
		r := mux.NewRouter()
		g := grpc.NewServer()


		ServerSingleton = &Server{
			router: r,
			svc:    service.TasksManagerSingleton,
			gsrv:   g,
			srv: &http.Server{
				Handler: r,
				Addr: ":" + strconv.Itoa(myconfig.ConfigSingleton.HTTPPort),
				WriteTimeout: time.Duration(myconfig.ConfigSingleton.HTTPWriteTimeout) * time.Millisecond ,
				ReadTimeout: time.Duration(myconfig.ConfigSingleton.HTTPReadTimeout) * time.Millisecond ,
			},
			ep: endpoints.EndpointsSingleton,
		}
	})
	return ServerSingleton
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

	endpoints.NewEndpoints()

	service.NewTasksManager(distribution.EtcdSingleton, mymongo.LogSingleton)

	newServerSingleton()

	gln, err := net.Listen("tcp", ":" + strconv.Itoa(myconfig.ConfigSingleton.GRPCPort))
	if err != nil {
		return
	}

	defer gln.Close()
	reflection.Register(ServerSingleton.gsrv)

	return
}