package service

import "context"

// interface redis cache
type IRedisCache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, expiratoinSeconds int) error
	Del(ctx context.Context, key string) error
	Incr(ctx context.Context, key string) (int64, error)
	Decr(ctx context.Context, key string) (int64, error)
	Exists(ctx context.Context, key string) (bool, error)
	WithDistributedLock(ctx context.Context, key string, expirationSeconds int, fn func(ctx context.Context) error) error
}

var vRedisCache IRedisCache

// Init singleton
func InitRedisCache(redis IRedisCache) {
	if vRedisCache != nil {
		panic("redis cache already initialized")
	}
	vRedisCache = redis
}

// GetRedisCache get redis cache
func GetRedisCache() IRedisCache {
	if vRedisCache == nil {
		panic("redis cache not initialized")
	}
	return vRedisCache
}
