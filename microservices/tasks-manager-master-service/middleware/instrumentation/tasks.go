package instrumentation

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"time"
)

func (mw *ImplInstrumentation) InstrumentationMiddleware(protocol, method string) endpoint.Middleware  {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				lvs := []string{"Protocol", protocol, "Method", method, "Error", fmt.Sprint(err != nil)}
				mw.requestCount.With(lvs...).Add(1)
				mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
			}(time.Now())
			return next(ctx, request)
		}
	}
}