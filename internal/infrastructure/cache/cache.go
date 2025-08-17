package cache

import (
	"context"
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
func (r *cacheRepository) SetView() error {
	key := "views"
	return r.cache.Incr(r.ctx, key).Err()
}
func (r *cacheRepository) GetView() error {
	key := "views"
	return r.cache.Get(r.ctx, key).Err()
}
func (r *cacheRepository) DelView() error {
	key := "views"
	return r.cache.Del(r.ctx, key).Err()
}
