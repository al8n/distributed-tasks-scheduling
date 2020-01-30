package encoder

import (
	"context"
	"sync"
)

type encoder interface {
	EncodeSaveOneTaskResponse(_ context.Context, response interface{}) (interface{}, error)
	EncodeDeleteOneTaskResponse(_ context.Context, response interface{}) (interface{}, error)
	EncodeGetOneTaskResponse(_ context.Context, response interface{}) (interface{}, error)
	EncodeGetAllTasksResponse(_ context.Context, response interface{}) (interface{}, error)
}

type ImplEncoder struct {}

var (
	EncoderSingleton *ImplEncoder
	once sync.Once
)

func NewEncoder() *ImplEncoder  {
	once.Do(func() {
		EncoderSingleton = &ImplEncoder{}
	})
	return EncoderSingleton
}

