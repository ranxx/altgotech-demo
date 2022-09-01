package model

import (
	"github.com/ranxx/altgotech-demo/pkg/db"
)

// User ...
type User struct {
	db.Base

	Email  string `gorm:"column:email;unique;type:varchar(64)"`
	Passwd string `gorm:"column:passwd;type:varchar(256)"`
}

// TableName ...
func (*User) TableName() string {
	return "user"
}
