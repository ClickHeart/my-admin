package main

import (
	"naive-backend/bootstrap"

	"github.com/donnie4w/go-logger/logger"
)

func main() {
	level := "debug"
	bootstrap.LogInit(level)
	logger.Debug("this is a debug message")
	logger.Info("this is an info message")
	logger.Warn("this is a warning message")
	logger.Error("this is an error message")
	logger.Fatal("this is a fatal message")
}
