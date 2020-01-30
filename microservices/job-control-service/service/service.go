package service

import (
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/entities"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/repositories/distribution"
	mymongo "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/repositories/mongo"
	"sync"
)

type TasksManagerService interface {
	SaveOneTask(name string) error
	DeleteOneTask(name string) error
	GetOneTask(name string) (entities.Task, error)
	GetAllTasks() ([]entities.Task, error)
}

type TasksManagerImpl struct {
	etcd *distribution.EtcdDB
	mongo *mymongo.LogDB
}

var (
	TasksManagerSingleton *TasksManagerImpl
	once sync.Once
)

func NewTasksManager(etcd *distribution.EtcdDB, mongo *mymongo.LogDB) *TasksManagerImpl {
	once.Do(func() {
		TasksManagerSingleton = &TasksManagerImpl{
			etcd:  etcd,
			mongo: mongo,
		}
	})
	return TasksManagerSingleton
}
