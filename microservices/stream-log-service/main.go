package main

import (
	sum "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/stream-log-service/grpc"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/stream-log-service/handler"
	"github.com/go-log/log"
	"github.com/micro/go-micro"
)

func main()  {
	// 创建服务

	srv := micro.NewService(
		micro.Name("go.micro.learning.srv.sum"),
	)

	srv.Init(micro.BeforeStart(func() error {
		log.Log("what are you doing")
		return nil
	}))
	//srv := server.NewServer(
	//	server.Name("go.micro.learning.srv.sum"),
	//)
	//// 初始化
	//srv.Init(
	//
	//	server.St
	//)

	// 挂载接口
	_ = sum.RegisterSumHandler(srv.Server(), handler.Handler())

	if err := srv.Run(); err != nil {
		panic(err)
	}
}