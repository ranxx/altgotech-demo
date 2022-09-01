package config

import (
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"os"
)

// Config 配置
type Config struct {
	Jwttoken string `json:"Jwttoken" yaml:"Jwttoken"`
	Mysql    Mysql  `json:"Mysql" yaml:"Mysql"`
}

type Mysql struct {
	Dsn string `json:"dsn" yaml:"dsn"`
}

var conf *Config = new(Config)

// GetConf ...
func GetConf() *Config {
	return conf
}

// GetLogio ...
func GetLogio() io.Writer {
	return os.Stdout
}

var (
	configPath = flag.String("config", "./config/config.json", "config file path")
)

func Init() {
	flag.Parse()
	data, err := ioutil.ReadFile(*configPath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, conf)
	if err != nil {
		panic(err)
	}
}
