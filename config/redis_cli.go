package config

import (
	"github.com/go-redis/redis/v9"
)

// var REDIS_ADDR = os.Getenv("REDIST_ADDR")
// var REDIS_DB = os.Getenv("REDIST_DB")
// var REDIS_USERNAME = os.Getenv("REDIST_USERNAME")
// var REDIS_PASSWORD = os.Getenv("REDIST_PASSWORD")

var REDIS_ADDR = "localhost:6379"
var REDIS_DB = 0
var REDIS_USERNAME = "havis"
var REDIS_PASSWORD = "secret"

func RedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     REDIS_ADDR,
		DB:       REDIS_DB,
		Username: REDIS_USERNAME,
		Password: REDIS_PASSWORD,
	})

	return rdb
}
