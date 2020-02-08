package entities

import (
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/conf"
	"strings"
)

func ExtractWorkerIP(regKey string) string  {
	return strings.TrimPrefix(regKey, conf.TASK_WORKER_DIR)
}
