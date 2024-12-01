package api

import "github.com/gin-gonic/gin"

type auth struct{}

var Auth = &auth{}

func (auth) Login(c *gin.Context) {

	Resp.Succ(c, "login success")
}
