package cache

import (
	"context"
	"fmt"
	"time"
)

// interface handle cache for auth
type (
	ICacheService interface {
		Get(ctx context.Context, key string) (interface{}, error)
		Set(ctx context.Context, key string, value interface{}) error
		SetWithTTL(ctx context.Context, key string, value interface{}, ttl time.Duration) error
		Del(ctx context.Context, key string) error
	}
)

// ########################################

var (
	vLocalCache       ICacheService
	vDistributedCache ICacheService
)

// ########################################

// get local cache instance
func GetLocalCacheService() (ICacheService, error) {
	if vLocalCache == nil {
		return nil, fmt.Errorf("local cache service is not initialized")
	}
	return vLocalCache, nil
}

// set local cache instance
func SetLocalCacheService(cache ICacheService) {
	vLocalCache = cache
}

// ########################################
// get distributed cache instance
func GetDistributedCacheService() (ICacheService, error) {
	if vDistributedCache == nil {
		return nil, fmt.Errorf("distributed cache service is not initialized")
	}
	return vDistributedCache, nil
}

// set distributed cache instance
func SetDistributedCacheService(cache ICacheService) {
	vDistributedCache = cache
}

// ########################################
