package config

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewRedis(ctx context.Context) *redis.Client {
	host := os.Getenv("REDIS_HOST")
	password := os.Getenv("REDIS_PASSWORD")
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		db = 0
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:            host,
		Password:        password,
		DB:              db,
		PoolSize:        5,
		MinIdleConns:    5,
		ConnMaxIdleTime: time.Minute * 5,
		ConnMaxLifetime: time.Minute * 30,
	})
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("failed connect to redis: %s", err)
	}
	return rdb
}
