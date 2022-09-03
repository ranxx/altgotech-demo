package migrate

import (
	"github.com/ranxx/altgotech-demo/pkg/db"
	amodel "github.com/ranxx/altgotech-demo/service/article/model"
	smodel "github.com/ranxx/altgotech-demo/service/space/model"
	umodel "github.com/ranxx/altgotech-demo/service/user/model"
)

// InitAutoMigrate 迁移表
func InitAutoMigrate() {
	mysql := db.Mysql()

	mysql.AutoMigrate(&umodel.User{})
	mysql.AutoMigrate(&amodel.Article{})
	mysql.AutoMigrate(&amodel.Comment{})
	mysql.AutoMigrate(&amodel.Like{})
	mysql.AutoMigrate(&smodel.Space{})
	mysql.AutoMigrate(&smodel.SpaceUser{})
}
