package decoder

import (
	"context"
	"net/http"
	"sync"
)

type HTTPDecoder interface {
	DecodeSaveOneTaskRequest(_ context.Context, r *http.Request) (interface{}, error)
	DecodeDeleteOneTaskRequest(_ context.Context, r *http.Request) (interface{}, error)
	DecodeGetOneTaskRequest(_ context.Context, r *http.Request) (interface{}, error)
	DecodeGetAllTasksRequest(_ context.Context, r *http.Request) (interface{}, error)
}

type ImplDecoder struct {}

var (
	SgtHTTPDecoder *ImplDecoder
	once sync.Once
)


func InitHTTPDecoder() *ImplDecoder  {
	once.Do(func() {
		SgtHTTPDecoder = &ImplDecoder{}
	})
	return SgtHTTPDecoder
}
