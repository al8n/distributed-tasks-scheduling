package logging

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"time"
)

func (mw *ImplLoggingMiddleware) LogMiddleware(protocol, method string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				SgtLogging.logger.Log("Protocol", protocol, "Method",method, "Error", fmt.Sprint(err != nil), "took", time.Since(begin))
			}(time.Now())

			return next(ctx, request)
		}
	}
}