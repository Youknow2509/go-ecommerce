package distributed

import (
	"context"
	"time"

	"github.com/Youknow2509/go-ecommerce/internal/user/domain/cache"
	"github.com/redis/go-redis/v9"
)

// Redis distributed cache structure
type RedisDistributedCache struct {
	client *redis.Client
}

// ########################################

// Del implements cache.ICacheService.
func (r *RedisDistributedCache) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

// Get implements cache.ICacheService.
func (r *RedisDistributedCache) Get(ctx context.Context, key string) (interface{}, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return val, nil
		}
		return val, err
	}
	return val, nil
}

// Set implements cache.ICacheService.
func (r *RedisDistributedCache) Set(ctx context.Context, key string, value interface{}) error {
	return r.client.Set(ctx, key, value, 0).Err()
}

// SetWithTTL implements cache.ICacheService.
func (r *RedisDistributedCache) SetWithTTL(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return r.client.Set(ctx, key, value, ttl).Err()
}

// ########################################

// Create redis distributed cache instance
func NewRedisDistributedCache(client *redis.Client) cache.ICacheService {
	return &RedisDistributedCache{
		client: client,
	}
}
