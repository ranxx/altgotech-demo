package main

import (
	"fmt"
	"net/http"

	"github.com/ranxx/altgotech-demo/config"
	"github.com/ranxx/altgotech-demo/pkg/db"
	"github.com/ranxx/altgotech-demo/router"
)

func init() {
	// 初始化 config
	config.Init()
	db.InitMysql(config.GetConf().Mysql.Dsn)
}

func main() {
	fmt.Println(http.ListenAndServe(":9123", router.Router()))

	fmt.Println("hello")
}
