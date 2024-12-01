package main

import (
	"fmt"
	"my-admin/internal/global"
	"my-admin/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.Default()
	router.Init(gin)
	gin.Run(fmt.Sprintf(":%d", global.Cfg.Server.Gin.Port))
}
