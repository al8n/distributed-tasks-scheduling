package encoder

import (
	"context"
	"encoding/json"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-master-service/responses"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/utils"
	"net/http"
)

func (ec *ImplEncoder) EncodeGetLogsResponse(_ context.Context, w http.ResponseWriter, response interface{}) error  {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	res, ok := response.(responses.GetLogsResponse)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return utils.EncoderAssertError
	}

	str, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write(str)
	return nil
}
