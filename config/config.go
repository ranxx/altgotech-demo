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
	Addr     string `json:"addr" yaml:"addr"`
	Jwttoken string `json:"jwttoken" yaml:"jwttoken"`
	Mysql    Mysql  `json:"mysql" yaml:"mysql"`
	Redis    Redis  `json:"redis" yaml:"redis"`
}

// Mysql ...
type Mysql struct {
	Dsn string `json:"dsn" yaml:"dsn"`
}

// Redis ...
type Redis struct {
	Host   string `json:"host" yaml:"host"`
	Passwd string `json:"passwd" yaml:"passwd"`
	DB     int    `json:"db" yaml:"db"`
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

// Init ...
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
