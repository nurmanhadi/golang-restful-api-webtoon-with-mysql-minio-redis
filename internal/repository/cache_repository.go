package repository

import "time"

type CacheRepository interface {
	Set(key string, value interface{}, exp time.Duration) error
	Get(key string) error
	Remove(key string) error
}
