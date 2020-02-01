package decoder

import (
	"context"
	"sync"
)

type GRPCDecoder interface {
	DecodeSaveOneTaskRequest(ctx context.Context, grpcReq interface{}) (interface{}, error)
	DecodeDeleteOneTaskRequest(ctx context.Context, grpcReq interface{}) (interface{}, error)
	DecodeGetAllTasksRequest(ctx context.Context, grpcReq interface{}) (interface{}, error)
	DecodeGetOneTaskRequest(ctx context.Context, grpcReq interface{}) (interface{}, error)
	DecodeKillOneTaskRequest(ctx context.Context, grpcReq interface{}) (interface{}, error)
}

type ImplDecoder struct {}

var (
	SgtGRPCDecoder *ImplDecoder
	once sync.Once
)

func InitGRPCDecoder() *ImplDecoder  {
	once.Do(func() {
		SgtGRPCDecoder = &ImplDecoder{}
	})
	return SgtGRPCDecoder
}