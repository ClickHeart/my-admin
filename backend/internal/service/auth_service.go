package service

import (
	"crypto/md5"
	"fmt"
	"my-admin/internal/model"
	"my-admin/internal/vo"
)

var Auth = &auth{}

type auth struct {
}

func (auth) Login(req vo.LoginReq) (token string, err error) {
	info := &model.User{
		Username: req.Username,
		Password: fmt.Sprintf("%x", md5.Sum([]byte(req.Password))),
	}
	var user model.User
	user, err = info.GetUser()
	if err != nil {
		return
	}
	token = fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%d", user.ID))))
	return
}
