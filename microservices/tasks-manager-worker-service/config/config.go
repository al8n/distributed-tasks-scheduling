package config

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

type Config struct {
	EtcdEndpoints []string `json:"etcd_endpoints"`
	EtcdTimeout int `json:"etcd_timeout"`
	TaskEventCapacity int `json:"task_event_capacity"`
	EtcdLockRandTime int `json:"etcd_lock_rand_time"`

	MongoEndpoint string `json:"mongo_endpoint"`
	MongoDatabase string `json:"mongo_database"`
	MongoCollection string `json:"mongo_collection"`
	MongoTimeout int `json:"mongo_timeout"`

	LogBatchSize int `json:"log_batch_size"`
	LogCommitTimeout int `json:"log_commit_timeout"`
	LogChanSize int `json:"log_chan_size"`
	LogCommitChanSize int `json:"log_commit_chan_size"`

	ShellCommand string `json:"shell_command"`
}

var (
	ConfigSingleton *Config
	once sync.Once
)

func newConfigSingleton(conf *Config) *Config  {
	once.Do(func() {
		ConfigSingleton = conf
	})

	return ConfigSingleton
}

func InitConfig(filename string) (err error) {
	var (
		content []byte
		conf Config
	)

	// 1. 把配置文件读进来
	if content, err = ioutil.ReadFile(filename); err != nil {
		return
	}

	// 2. 做Json反序列化
	if err = json.Unmarshal(content, &conf); err != nil {
		return
	}

	// 3. 赋值单例
	newConfigSingleton(&conf)
	return
}

