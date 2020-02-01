package encoder

import (
	"context"
	"net/http"
	"sync"
)

type HTTPEncoder interface {
	EncodeSaveOneTaskResponse(_ context.Context, w http.ResponseWriter, response interface{}) error
	EncodeDeleteOneTaskResponse(_ context.Context, w http.ResponseWriter, response interface{}) error
	EncodeKillOneTaskResponse(_ context.Context, w http.ResponseWriter, response interface{}) error
	EncodeGetOneTaskResponse(_ context.Context, w http.ResponseWriter, response interface{}) error
	EncodeGetAllTasksResponse(_ context.Context, w http.ResponseWriter, response interface{}) error
}

type ImplEncoder struct {}


var (
	SgtHTTPEncoder *ImplEncoder
	once sync.Once
)


func InitHTTPEncoder() *ImplEncoder  {
	once.Do(func() {
		SgtHTTPEncoder = &ImplEncoder{}
	})
	return SgtHTTPEncoder
}