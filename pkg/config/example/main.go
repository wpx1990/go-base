package main

import (
	"github.com/wpx1990/go-base/pkg/config"
	"github.com/wpx1990/go-base/pkg/log"
)

type Config struct {
	LogLevel string `yaml:"LogLevel"` // 日志级别
	IPAddr   string `yaml:"IPAddr"`   // 地址
	Port     uint32 `yaml:"Port"`     // 端口
}

func main() {

	// 初始化日志打印
	log.InitLogger("debug", "")
	defer log.ReleaseLogger()
	log.Info("Succeed to init logger.")

	var conf Config

	// 读取配置文件
	err := config.GetConfig([]string{"./example.yaml"}, &conf)
	if err != nil {
		log.Error("Failed to get config, err(%v).", err)
		return
	}
	log.Info("Succeed to init config(%+v).", conf)
}
