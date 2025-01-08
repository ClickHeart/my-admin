package router

import (
	"my-admin/internal/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Init(g *gin.Engine) {
	apiV1 := g.Group("v1")
	apiV1.Use(sessions.Sessions("mysession", cookie.NewStore([]byte("captch"))))
	apiV1.Use(middleware.Cors())

	WithAuthRouter(apiV1)
}
