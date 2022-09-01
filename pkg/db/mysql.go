package db

import (
	"time"

	"github.com/ranxx/altgotech-demo/pkg/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlDB *gorm.DB

func InitMysql(host string) {
	var err error
	mysqlDB, err = gorm.Open(mysql.Open(host), &gorm.Config{})
	if err != nil {
		log.Logrus().Panic("数据库链接错误")
	}
	sqlDB, _ := mysqlDB.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(59 * time.Second)
}

func Mysql() *gorm.DB {
	return mysqlDB
}

// Pager 页面
func Pager(in *gorm.DB, offset, limit int64) *gorm.DB {
	ret := in
	if limit >= 0 {
		ret = ret.Limit(int(limit))
	}
	if offset > 0 && limit >= 0 {
		ret = ret.Offset(int((offset - 1) * limit))
	}
	return ret
}
