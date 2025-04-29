package impl

import (
	"context"
	"fmt"
	"time"

	"github.com/Youknow2509/go-ecommerce/internal/service"
	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
)

// struct redis cache
type sRedisCache struct {
	client *redis.Client
	locker *redislock.Client
}

// WithDistributedLock implements service.IRedisCache.
func (s *sRedisCache) WithDistributedLock(ctx context.Context, key string, expirationSeconds int, fn func(ctx context.Context) error) error {
	timeTll := time.Second * time.Duration(expirationSeconds)
	// linearBackoff time
	// linearBackoff := redislock.LinearBackoff(500 * time.Millisecond)
	lock, err := s.locker.Obtain(
		ctx,
		key,
		timeTll,
		&redislock.Options{
			// RetryStrategy: linearBackoff,
		},
	)
	if err == redislock.ErrNotObtained {
		return fmt.Errorf("could not obtain lock for key: %s", key)
	} else if err != nil {
		return fmt.Errorf("failed to obtain lock: %w", err)
	}
	defer lock.Release(ctx)

	return fn(ctx)
}

// Decr implements service.IRedisCache.
func (s *sRedisCache) Decr(ctx context.Context, key string) (int64, error) {
	val, err := s.client.Decr(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return val, nil
}

// Del implements service.IRedisCache.
func (s *sRedisCache) Del(ctx context.Context, key string) error {
	err := s.client.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

// Exists implements service.IRedisCache.
func (s *sRedisCache) Exists(ctx context.Context, key string) (bool, error) {
	val, err := s.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	if val == 0 {
		return false, nil
	}
	return true, nil
}

// Get implements service.IRedisCache.
func (s *sRedisCache) Get(ctx context.Context, key string) (string, error) {
	val, err := s.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return val, nil
		}
		return val, err
	}
	return val, nil
}

// Incr implements service.IRedisCache.
func (s *sRedisCache) Incr(ctx context.Context, key string) (int64, error) {
	val, err := s.client.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return val, nil
}

// Set implements service.IRedisCache.
func (s *sRedisCache) Set(ctx context.Context, key string, value interface{}, expiratoinSeconds int) error {
	ex := time.Second * time.Duration(expiratoinSeconds)
	err := s.client.Set(ctx, key, value, ex).Err()
	if err != nil {
		return err
	}
	return nil
}

// NewRedisCache create a new redis cache
func NewRedisCache(client *redis.Client) service.IRedisCache {
	return &sRedisCache{
		client: client,
		locker: redislock.New(client),
	}
}
