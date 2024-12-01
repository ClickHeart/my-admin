package router

import (
	"my-admin/internal/api"

	"github.com/gin-gonic/gin"
)

func NewAuthRouter(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	auth.GET("/login", api.Auth.Login)
}
