package cache

import (
	"context"
	"strconv"
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
func (r *cacheRepository) GetView() (int, error) {
	var total int
	key := "views"
	v, err := r.cache.Get(r.ctx, key).Result()
	if err != nil {
		return 0, nil
	}
	total, err = strconv.Atoi(v)
	if err != nil {
		return 0, err
	}
	return total, nil
}
func (r *cacheRepository) DelView() error {
	key := "views"
	return r.cache.Del(r.ctx, key).Err()
}
