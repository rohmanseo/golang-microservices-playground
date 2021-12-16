package memory

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

func NewCacheDataStore(redisDb *redis.Client, ctx *context.Context) ICacheDataStore {
	return &cacheDataStoreImpl{
		redisDb: redisDb,
		ctx:     ctx,
	}
}

type cacheDataStoreImpl struct {
	redisDb *redis.Client
	ctx     *context.Context
}

func (c *cacheDataStoreImpl) AddUser() {
	result := c.redisDb.Get(*c.ctx, "user_count").Val()
	if result == "0" || result == "" {
		c.redisDb.Set(*c.ctx, "user_count", 1, 24*time.Hour)
	} else {
		resInt, _ := strconv.Atoi(result)
		c.redisDb.Set(*c.ctx, "user_count", resInt+1, 24*time.Hour)
	}
}

func (c *cacheDataStoreImpl) GetStatistic() int {
	result := c.redisDb.Get(*c.ctx, "user_count").Val()

	if result == "0" || result == "" {
		return 0
	}
	res, _ := strconv.Atoi(result)
	return res
}
