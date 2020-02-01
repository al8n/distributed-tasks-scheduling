package instrumentation

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/metrics"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"sync"
)

type Instrumentation interface {
	TaskInstrumentationMiddleware(protocol, method string) endpoint.Middleware
}

type ImplInstrumentation struct {
	requestCount metrics.Counter
	requestLatency metrics.Histogram
	countResult metrics.Histogram
}


var (
	SgtInstrumentation *ImplInstrumentation
	once sync.Once
)

func InitInstrumentationMiddleware()  *ImplInstrumentation {
	fieldKeys := []string{"Protocol", "Method", "Error"}

	rc := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "Task_Manager",
		Subsystem: "Task_Services",
		Name: "Request_Count",
		Help: "Number of requests received.",
	}, fieldKeys)

	rl := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "Task_Manager",
		Subsystem: "Task_Services",
		Name:      "Request_Latency_Microseconds",
		Help:      "Total duration of requests in microseconds.",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},fieldKeys)

	cr := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "Task_Manager",
		Subsystem: "Task_Services",
		Name:      "Count_Result",
		Help:      "The result of each count method.",
	}, []string{})

	once.Do(func() {
		SgtInstrumentation = &ImplInstrumentation{
			requestCount:   rc,
			requestLatency: rl,
			countResult:    cr,
		}
	})
	return SgtInstrumentation
}

