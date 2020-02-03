package services

import (
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/entities"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/requests"
	"sync"
)

type Service interface {
	SaveOneTask(task entities.Task) (oldTask *entities.Task, err error)
	DeleteOneTask(name string) (oldTask *entities.Task, err error)
	KillOneTask(name string) ( err error)
	GetOneTask(name string) (*entities.Task, error)
	GetAllTasks() ([]*entities.Task, error)

	GetLogs(req requests.GetLogsRequest) ([]*entities.TaskLog, error)
}

type ImplService struct {}

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