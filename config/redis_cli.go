package config

import (
	"os"

	"github.com/go-redis/redis/v9"
)

var REDIS_ADDR = os.Getenv("REDIS_ADDR")
var REDIS_DB = 0
var REDIS_USERNAME = os.Getenv("REDIS_USERNAME")
var REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")

func RedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     REDIS_ADDR,
		DB:       REDIS_DB,
		Username: REDIS_USERNAME,
		Password: REDIS_PASSWORD,
	})

	return rdb
}
