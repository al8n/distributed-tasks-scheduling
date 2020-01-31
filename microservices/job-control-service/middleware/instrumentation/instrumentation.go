package instrumentation

import (
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/metrics"
)

type Instrumentation interface {
	NewMetrics(requestCount metrics.Counter, requestLatency metrics.Histogram) endpoint.Middleware
}

type ImplInstrumentation struct {
	service.Service
	requestCount metrics.Counter
	requestLatency metrics.Histogram
	countResult metrics.Histogram
}


