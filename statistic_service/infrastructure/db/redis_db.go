package db

import (
	"github.com/go-redis/redis/v8"
	"github.com/rohmanseo/golang-clean-arch/config"
	"strconv"
)

func NewRedisDb(config config.IConfig) *redis.Client {
	db, _ := strconv.Atoi(config.Get("REDIS_DB"))
	opt := &redis.Options{
		Addr:     config.Get("REDIS_ADDRESS"),
		Password: config.Get("REDIS_PASSWORD"),
		DB:       db, // use default DB
	}

	return redis.NewClient(opt)
}
