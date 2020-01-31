package main

import (
	"flag"
	"fmt"
	myconfig "github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/config"
	"github.com/ALiuGuanyan/distributed-task-scheduling/microservices/job-control-service/server"
	"log"
	"runtime"
)

var (
	confFile string // 配置文件路径
)

// 解析命令行参数
func initArgs() {
	// master -config ./master.json -xxx 123 -yyy ddd
	// master -h
	flag.StringVar(&confFile, "config", "../conf/config.json", "指定config.json")
	flag.Parse()
}

// 初始化线程数量
func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}


func main () {
	var (
		err error
	)

	// 初始化命令行参数
	initArgs()

	// 初始化线程
	initEnv()

	// 加载配置
	if err = myconfig.InitConfig(confFile); err != nil {
		log.Println(err)
	}

	fmt.Println("Serve")
	err = server.InitServer()
	if err != nil {
		log.Println(err)
	}
}
