package main

import (
	"fmt"
	"my-admin/internal/global"

	"github.com/donnie4w/go-logger/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.Default()
	logger.Info("this is an info message")
	logger.Warn("this is a warning message")
	logger.Error("this is an error message")
	logger.Fatal("this is a fatal message")
	gin.Run(fmt.Sprintf(":%d", global.Cfg.Server.Gin.Port))
}
