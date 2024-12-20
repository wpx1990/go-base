package main

import (
	"github.com/wpx1990/go-base/pkg/log"
)

func main() {

	// 初始化日志打印
	log.InitLogger("debug", "")
	defer log.ReleaseLogger()
	log.Info("Succeed to init logger.")

	log.Debug("This is a debug log.")
	log.Info("This is a info log.")
	log.Warn("This is a warn log.")
	log.Error("This is a error log.")

	// 提升日志级别
	log.IncreaseLogLevel("info")

	log.Debug("This is a debug log.")
	log.Info("This is a info log.")
	log.Warn("This is a warn log.")
	log.Error("This is a error log.")
	log.Panic("This is a panic log.")
}
