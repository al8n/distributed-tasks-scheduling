package transport

import (
	logsTransport "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/grpc/transport/logs_transport"
	tasksTransport "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/grpc/transport/tasks_transport"
)

type Transport struct {
	Logs *logsTransport.LogsTransport
	Tasks *tasksTransport.TasksTransport
}


