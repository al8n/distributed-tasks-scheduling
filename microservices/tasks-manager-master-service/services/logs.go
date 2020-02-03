package services

import (
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/entities"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/repositories/mongo"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/requests"
)

func (svc *ImplService) GetLogs(req requests.GetLogsRequest) ([]*entities.TaskLog, error) {
	return mongo.SgtLogDB.GetLogs(req)
}
