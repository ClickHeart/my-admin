package router

import (
	"my-admin/internal/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Init(g *gin.Engine) {
	r := g.Group("v1")
	r.Use(sessions.Sessions("mysession", cookie.NewStore([]byte("captch"))))
	r.Use(middleware.Cors())

	NewAuthRouter(r)
}
