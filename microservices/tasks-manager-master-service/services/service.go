package services

import (
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/entities"
	mygrpctransport "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/grpc/transport"
	"sync"
)

type Service interface {
	SaveOneTask(task entities.Task) (oldTask *entities.Task, err error)
	DeleteOneTask(name string) (oldTask *entities.Task, err error)
	KillOneTask(name string) ( err error)
	GetOneTask(name string) (*entities.Task, error)
	GetAllTasks() ([]*entities.Task, error)
}

type ImplService struct {
	//instrumentation.Instrumentation
	tp mygrpctransport.ImplTasksTransport
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