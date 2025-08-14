package cache

import (
	"context"
	"time"
	"welltoon/internal/repository"

	"github.com/redis/go-redis/v9"
)

type cacheRepository struct {
	ctx   context.Context
	cache *redis.Client
}

func NewCache(ctx context.Context, cache *redis.Client) repository.CacheRepository {
	return &cacheRepository{ctx: ctx, cache: cache}
}
func (r *cacheRepository) Set(key string, value interface{}, exp time.Duration) error {
	return r.cache.SetEx(r.ctx, key, value, exp).Err()
}
func (r *cacheRepository) Get(key string) error {
	return r.cache.Get(r.ctx, key).Err()
}
func (r *cacheRepository) Remove(key string) error {
	return r.cache.Del(r.ctx, key).Err()
}
