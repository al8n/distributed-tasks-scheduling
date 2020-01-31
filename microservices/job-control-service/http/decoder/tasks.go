package decoder

import (
	"context"
	"encoding/json"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/requests"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/utils/appErrors"
	"github.com/gorilla/mux"
	"net/http"
)

func (dc *ImplDecoder) DecodeSaveOneTaskRequest(_ context.Context, r *http.Request) (i interface{}, err error) {
	var request requests.SaveOneTaskRequest
	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		return
	}

	return request, nil
}

func (dc *ImplDecoder) DecodeDeleteOneTaskRequest(_ context.Context, r *http.Request) (i interface{}, err error)  {
	var request requests.DeleteOneTaskRequest
	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		return
	}

	return request, nil
}

func (dc *ImplDecoder) DecodeGetOneTaskRequest(_ context.Context, r *http.Request) (i interface{}, err error)  {
	vars := mux.Vars(r)
	bd, ok := vars["name"]
	if !ok {
		return nil, appErrors.GetRequestParamsError
	}

	return requests.GetOneTaskRequest{
		Name: bd,
	}, nil
}

func (dc *ImplDecoder) DecodeGetAllTasksRequest(_ context.Context, r *http.Request) (i interface{}, err error)  {
	return requests.GetAllTasksRequest{}, nil
}
