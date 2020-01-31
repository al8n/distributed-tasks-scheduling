package service

import (
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/entities"
	mygrpctransport "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/grpc/transport"
	"sync"
)

type Service interface {
	SaveOneTask(task entities.Task) (oldTask *entities.Task, err error)
	DeleteOneTask(name string) (oldTask *entities.Task, err error)
	GetOneTask(name string) (*entities.Task, error)
	GetAllTasks() ([]*entities.Task, error)
}

type ImplService struct {
	//instrumentation.Instrumentation
	mygrpctransport.Transports
}

var (
	SgtService *ImplService
	once sync.Once
)

func InitService() *ImplService {
	once.Do(func() {
		SgtService = &ImplService{}
	})
	return SgtService
}