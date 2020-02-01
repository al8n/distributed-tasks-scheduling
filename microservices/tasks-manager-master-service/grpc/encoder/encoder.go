package encoder

import (
	"context"
	"sync"
)

type GRPCEncoder interface {
	EncodeSaveOneTaskResponse(_ context.Context, response interface{}) (interface{}, error)
	EncodeDeleteOneTaskResponse(_ context.Context, response interface{}) (interface{}, error)
	EncodeKillOneTaskResponse(_ context.Context, response interface{}) (interface{}, error)
	EncodeGetOneTaskResponse(_ context.Context, response interface{}) (interface{}, error)
	EncodeGetAllTasksResponse(_ context.Context, response interface{}) (interface{}, error)
}

type ImplEncoder struct {}

var (
	SgtGRPCEncoder *ImplEncoder
	once sync.Once
)

func InitGRPCEncoder() *ImplEncoder  {
	once.Do(func() {
		SgtGRPCEncoder = &ImplEncoder{}
	})
	return SgtGRPCEncoder
}

