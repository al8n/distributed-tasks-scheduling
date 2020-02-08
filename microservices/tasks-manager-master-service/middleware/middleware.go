package middleware

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"sync"
)

type middleware interface {
	LoggingMiddleware(logger log.Logger) endpoint.Middleware
}

type Middleware struct {

}


var (
	MiddlewareSingleton *Middleware
	once sync.Once
)


