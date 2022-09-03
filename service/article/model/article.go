package model

import (
	"github.com/ranxx/altgotech-demo/pkg/db"
)

// Article 文章
type Article struct {
	db.Base

	Title     string `gorm:"column:title;type:varchar(256);comment:文章标题"`
	Content   string `gorm:"column:content;type:text;comment:文章内容"`
	Tags      string `gorm:"column:tags;type:varchar(256);comment:文章tags,多个用逗号隔开"`
	SpaceID   int    `gorm:"column:space_id;size:64;comment:哪个板块下面"`
	CreatoriD int    `gorm:"column:creator_id;size:64;comment:创建人id"`
	Top       bool   `gorm:"column:top;size:8;comment:置顶;default:false"`
}

// TableName ...
func (*Article) TableName() string {
	return "article"
}

// Comment 评论
type Comment struct {
	db.Base

	Content string `gorm:"column:content;type:text;comment:评论内容"`
	AID     int    `gorm:"column:aid;size:64;comment:文章id"`
	CID     int    `gorm:"column:cid;size:64;comment:回复评论时的id"`
	UID     int    `gorm:"column:uid;size:64;comment:创建人id"`
}

// TableName ...
func (*Comment) TableName() string {
	return "comment"
}

// Like 点赞
type Like struct {
	db.Base

	AID int `gorm:"column:aid;size:64;comment:文章id"`
	UID int `gorm:"column:uid;size:64;comment:创建人id"`
}

// TableName ...
func (*Like) TableName() string {
	return "like"
}

// Total ...
type Total struct {
	ID    int   `json:"id"`
	Total int64 `json:"total"`
}
