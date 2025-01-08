package model

import (
	"my-admin/internal/global"
	"time"
)

type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Enable     bool      `json:"enable"`
	CreateTime time.Time `json:"createTime" gorm:"column:createTime"`
	UpdateTime time.Time `json:"updateTime" gorm:"column:updateTime"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) GetUser() (user User, err error) {
	err = global.DB.Model(u).
		Where("username =? ", u.Username).
		Where("password=?", u.Password).
		Find(&user).Error
	return
}
