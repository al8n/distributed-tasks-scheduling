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

