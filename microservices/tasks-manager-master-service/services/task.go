package services

import (
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/entities"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/repositories/distribution"
)

func (svc *ImplService) SaveOneTask(task entities.Task) (oldTask *entities.Task, err error) {

	return distribution.SgtEtcd.SaveOneTask(&task)
}


func (svc *ImplService) DeleteOneTask(name string) (oldTask *entities.Task, err error) {
	return distribution.SgtEtcd.DeleteOneTask(name)
}


func (svc *ImplService) GetOneTask(name string) (task *entities.Task, err error) {
	return distribution.SgtEtcd.GetOneTask(name)
}


func (svc *ImplService) GetAllTasks() (tasks []*entities.Task, err error) {
	return distribution.SgtEtcd.GetAllTasks()
}

func (svc *ImplService) KillOneTask(name string) (err error)  {
	return distribution.SgtEtcd.KillTask(name)
}