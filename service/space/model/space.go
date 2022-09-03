package model

import (
	"github.com/ranxx/altgotech-demo/pkg/db"
)

// Space 板块
type Space struct {
	db.Base

	Name      string `gorm:"column:name;type:varchar(64);comment:板块名"`
	CreatorID int    `gorm:"column:creator_id;size:64;comment:创建人id"`
	AdminID   int    `gorm:"column:admin_id;size:64;default:-1;comment:管理员id"`
}

// TableName ...
func (*Space) TableName() string {
	return "space"
}

// SpaceUser 加入的板块
type SpaceUser struct {
	db.Base

	SID    int `gorm:"column:sid;size:64;comment:板块id"`
	UserID int `gorm:"column:user_id;size:64;comment:用户id"`
}

// TableName ...
func (*SpaceUser) TableName() string {
	return "space_user"
}
