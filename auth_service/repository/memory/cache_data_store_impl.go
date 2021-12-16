package memory

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/rohmanseo/golang-clean-arch/entity"
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

func (c *cacheDataStoreImpl) AddToken(token entity.Token) {
	key := fmt.Sprintf("token:%s", token.AccessToken)
	c.redisDb.Set(*c.ctx, key, 1, time.Hour*24)
}

func (c *cacheDataStoreImpl) RevokeToken(token entity.Token) (bool, error) {
	key := fmt.Sprintf("token:%s", token.AccessToken)
	c.redisDb.Set(*c.ctx, key, 0, 0)
	return true, nil
}

func (c *cacheDataStoreImpl) ValidateToken(token entity.Token) bool {
	key := fmt.Sprintf("token:%s", token.AccessToken)
	result := c.redisDb.Get(*c.ctx, key).Val()
	if result == "0" || result == "" {
		return false
	}
	return true
}
