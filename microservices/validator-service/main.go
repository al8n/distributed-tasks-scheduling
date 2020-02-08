package main

import (
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/entities"
	"net/http"
)

func main()  {
	// 过滤器
	filter := entities.NewFilter()

	// 注册拦截器
	filter.RegisterFilterUri("/check", Auth())

	// 启动服务


}

// 统一验证拦截器
func Auth(w http.ResponseWriter, r *http.Request) error {
	return nil
}