package decoder

import (
	"context"
	"encoding/json"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/requests"
	"net/http"
)

func (dc *ImplDecoder) DecodeSaveOneTaskRequest(_ context.Context, r *http.Request) (i interface{}, err error) {
	var request requests.SaveOneTaskRequest
	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		return
	}

	return request, nil
}

func (dc *ImplDecoder)  DecodeKillOneTaskRequest(_ context.Context, r *http.Request) (i interface{}, err error)  {
	var request requests.KillOneTaskRequest
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

	var request requests.DeleteOneTaskRequest
	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		return
	}

	return requests.GetOneTaskRequest{
		Name: request.Name,
	}, nil
}

func (dc *ImplDecoder) DecodeGetAllTasksRequest(_ context.Context, r *http.Request) (i interface{}, err error)  {
	return requests.GetAllTasksRequest{}, nil
}
