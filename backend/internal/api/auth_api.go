package api

import (
	"errors"
	"my-admin/internal/service"
	"my-admin/internal/vo"
	"my-admin/pkg/utils"
	"net/http"

	"github.com/donnie4w/go-logger/logger"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type auth struct{}

var Auth = &auth{}

func (auth) Captcha(c *gin.Context) {
	svg, code := utils.GenerateSVG(80, 40)
	session := sessions.Default(c)
	session.Set("captch", code)
	session.Save()
	// 设置 Content-Type 为 "image/svg+xml"
	c.Header("Content-Type", "image/svg+xml; charset=utf-8")
	// 返回验证码
	c.Data(http.StatusOK, "image/svg+xml", svg)
}

func (auth) Login(c *gin.Context) {
	var req vo.LoginReq
	err := c.Bind(&req)
	if err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	session := sessions.Default(c)
	if req.Captcha != session.Get("captch") {
		Resp.Err(c, 20001, "验证码不正确")
		return
	}

	token, err := service.Auth.Login(req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			Resp.Err(c, 20001, "账号或密码不正确")
		} else {
			Resp.Err(c, 20001, "服务器内部错误")
			logger.Errorf("GetUser err: %v", err)
		}
		return
	}

	Resp.Succ(c, vo.LoginRes{
		AccessToken: token,
	})
}
