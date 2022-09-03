package db

import (
	"github.com/go-redis/redis/v8"
	"github.com/ranxx/altgotech-demo/config"
)

var redisDB *redis.Client

// InitRedis init redis
func InitRedis(rc config.Redis) {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     rc.Host,
		Password: rc.Passwd,
		DB:       rc.DB,
	})
}

// Redis ...
func Redis() *redis.Client {
	return redisDB
}
