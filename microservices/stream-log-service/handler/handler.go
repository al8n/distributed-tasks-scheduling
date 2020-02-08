package handler

import (
	"context"
	sum "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/stream-log-service/grpc"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/stream-log-service/service"
)

type handler struct {

}

func  (h *handler) GetSum(ctx context.Context, req *sum.SumRequest, rep *sum.SumResponse) error {
	inputs := make([]int64, 0)

	var i int64 = 0

	for ; i <= req.Input; i++ {
		inputs = append(inputs, i)
	}

	rep.Output = service.GetSum(inputs...)
	return nil
}

func Handler() sum.SumHandler  {
	return &handler{}
}