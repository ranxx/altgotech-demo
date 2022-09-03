package main

import (
	"fmt"
	"net/http"

	"github.com/ranxx/altgotech-demo/config"
	"github.com/ranxx/altgotech-demo/pkg/db"
	"github.com/ranxx/altgotech-demo/router"
	"github.com/ranxx/altgotech-demo/service/migrate"
)

func init() {
	// 初始化 config
	config.Init()
	db.InitMysql(config.GetConf().Mysql)
	db.InitRedis(config.GetConf().Redis)
	migrate.InitAutoMigrate()

	db.SetMysqlDebug(true) // debug
}

func main() {
	fmt.Println(http.ListenAndServe(config.GetConf().Addr, router.Router()))
}
