package config

import (
	"github.com/go-redis/redis/v9"
)

// var REDIS_ADDR = os.Getenv("REDIS_ADDR")
// var REDIS_PASS = os.Getenv("REDIS_PASS")
// var REDIS_USERNAME = os.Getenv("REDIS_USERNAME")

var REDIS_ADDR = "localhost:6379"
var REDIS_PASS = "secret"
var REDIS_USERNAME = "havis"
var REDIS_DB = 0

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     REDIS_ADDR,
		DB:       REDIS_DB,
		Username: REDIS_USERNAME,
		Password: REDIS_PASS,
	})
	return rdb
}
