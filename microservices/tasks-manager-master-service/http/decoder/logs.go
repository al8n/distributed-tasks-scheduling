package decoder

import (
	"context"
	"encoding/json"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/requests"
	"net/http"
)

func (dc *ImplDecoder) DecodeGetLogsRequest(_ context.Context, r *http.Request) (i interface{}, err error)  {
	var request requests.GetLogsRequest
	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		return
	}

	return request, nil
}


