package model

type UserRoles struct {
	UserId int `gorm:"column:userId"`
	RoleId int `gorm:"column:roleId"`
}

func (UserRoles) TableName() string {
	return "user_roles"
}
