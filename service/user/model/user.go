package model

import (
	"github.com/ranxx/altgotech-demo/pkg/db"
)

// User ...
type User struct {
	db.Base

	Name   string `gorm:"column:name;unique;type:varchar(64);comment:用户名"`
	Email  string `gorm:"column:email;unique;type:varchar(64);comment:邮箱"`
	Passwd string `gorm:"column:passwd;type:varchar(512);comment:密码"`
}

// TableName ...
func (*User) TableName() string {
	return "user"
}
