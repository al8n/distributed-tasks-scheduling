package logging

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"os"
	"sync"
)

type LoggingMiddleware interface {
	TaskLoggingMiddleware(protocol, method string) endpoint.Middleware
}

type ImplLoggingMiddleware struct {
	logger log.Logger
}

var (
	SgtLogging *ImplLoggingMiddleware
	once sync.Once
)

func InitLoggingMiddleware() *ImplLoggingMiddleware  {
	once.Do(func() {
		logger := log.NewLogfmtLogger(os.Stderr)
		SgtLogging = &ImplLoggingMiddleware{
			logger: logger,
		}
	})
	return SgtLogging
}







