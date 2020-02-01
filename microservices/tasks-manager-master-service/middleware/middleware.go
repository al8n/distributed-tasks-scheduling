package middleware

import (
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/tasks-manager-master-service/middleware/logging"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"sync"
)

type middleware interface {
	LoggingMiddleware(logger log.Logger) endpoint.Middleware
}

type ImplMiddleware struct {
	log *logging.ImplLoggingMiddleware
}


var (
	MiddlewareSingleton *ImplMiddleware
	once sync.Once
)

func NewMiddleware(logger *log.Logger) *ImplMiddleware {
	once.Do(func() {
		MiddlewareSingleton = &ImplMiddleware{
			log: logging.NewLoggingMiddleware(logger),
		}
	})
	return MiddlewareSingleton
}
