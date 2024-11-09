package main

import (
	"naive-backend/bootstrap"
	"naive-backend/global"

	"github.com/donnie4w/go-logger/logger"
)

func main() {
	if err := bootstrap.ConfigInit(); err != nil {
		return
	}
	bootstrap.LogInit(global.Cfg.Log)
	logger.Debug("this is a debug message")
	logger.Info("this is an info message")
	logger.Warn("this is a warning message")
	logger.Error("this is an error message")
	logger.Fatal("this is a fatal message")
}
