package decoder

import (
	"context"
	"sync"
)

type decoder interface {
	DecodeSaveOneTaskRequest(ctx context.Context, grpcReq interface{}) (interface{}, error)
	DecodeDeleteOneTaskRequest(ctx context.Context, grpcReq interface{}) (interface{}, error)
	DecodeGetAllTasksRequest(ctx context.Context, grpcReq interface{}) (interface{}, error)
	DecodeGetOneTaskRequest(ctx context.Context, grpcReq interface{}) (interface{}, error)
}

type ImplDecoder struct {}

var (
	DecoderSingleton *ImplDecoder
	once sync.Once
)

func NewDecoder() *ImplDecoder  {
	once.Do(func() {
		DecoderSingleton = &ImplDecoder{}
	})
	return DecoderSingleton
}