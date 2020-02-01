package responses

import (
	"encoding/json"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/entities"
)

type Response struct {
	Errno int32  `json:"error_number"`
	Message string `json:"msg"`
	Data interface{} `json:"data"`
}


type GetOneTaskResponse struct {
	Errno int32         `json:"error_number"`
	Message string      `json:"msg"`
	Task *entities.Task `json:"task"`
}

type GetAllTasksResponse struct {
	Errno int32            `json:"error_number"`
	Message string         `json:"msg"`
	Tasks []*entities.Task `json:"tasks"`
}

// 应答方法
func BuildResponse(errno int32, msg string, data interface{})(resp []byte, err error)  {
	var (
		response Response
	)

	response.Errno = errno
	response.Message = msg
	response.Data = data

	// 2. 序列化
	resp, err = json.Marshal(response)

	return
}

