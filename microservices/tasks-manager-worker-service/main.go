package main

import (
	"flag"
	myconfig "github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-worker-service/config"
	"github.com/ALiuGuanyan/distributed-tasks-scheduling/microservices/tasks-manager-worker-service/repositories/distribution"
	"log"
	"runtime"
	"time"
)

var (
	confFile string // 配置文件路径
)

// 解析命令行参数
func initArgs() {
	// master -config ./master.json -xxx 123 -yyy ddd
	// master -h
	flag.StringVar(&confFile, "config", "./config.json", "please specify config file")
	flag.Parse()
}

// 初始化线程数量
func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main()  {
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

	// 启动日志协程
	distribution.InitSink()

	// 初始化执行器
	distribution.InitExecutor()

	// 初始化调度器
	distribution.InitScheduler()

	// 初始化任务管理器
	if err = distribution.InitEtcd(); err != nil {
		log.Println(err)
		return
	}



	// 正常退出
	for {
		time.Sleep(1 * time.Second)
	}



}