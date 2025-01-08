package router

import (
	"my-admin/internal/api"

	"github.com/gin-gonic/gin"
)

func WithAuthRouter(r *gin.RouterGroup) *gin.RouterGroup {
	auth := r.Group("/auth")
	auth.GET("/login", api.Auth.Login)
	return auth
}
