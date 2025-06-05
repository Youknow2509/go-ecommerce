package local

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Youknow2509/go-ecommerce/internal/auth/domain/cache"
	"github.com/dgraph-io/ristretto/v2"
)

// Ristretto local cache structure
type RistrettoLocalCache struct {
	client *ristretto.Cache[string, string]
}
// ########################################

// Del implements cache.ICacheService.
func (r *RistrettoLocalCache) Del(ctx context.Context, key string) error {
	r.client.Del(key)
	return nil
}

// Get implements cache.ICacheService.
func (r *RistrettoLocalCache) Get(ctx context.Context, key string) (interface{}, error) {
	value, found := r.client.Get(key)
	if !found {
		return nil, nil
	}
	var data interface{}
	err := json.Unmarshal([]byte(value), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Set implements cache.ICacheService.
func (r *RistrettoLocalCache) Set(ctx context.Context, key string, value interface{}) error {
	jsonData, _ := json.Marshal(value)
	ok := r.client.Set(key, string(jsonData), 1)
	if !ok {
		return fmt.Errorf("failed to set cache")
	}
	return nil
}

// SetWithTTL implements cache.ICacheService.
func (r *RistrettoLocalCache) SetWithTTL(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	jsonData, _ := json.Marshal(value)
	ok := r.client.SetWithTTL(key, string(jsonData), 1, ttl)
	if !ok {
		return fmt.Errorf("failed to set cache with TTL")
	}
	return nil
}

// ########################################

// new instance of RistrettoLocalCache
func NewReistrettoLocalCache() cache.ICacheService {
	cache, _ := ristretto.NewCache(
		&ristretto.Config[string, string]{
			NumCounters: 1e7,     // number of keys to track frequency of (10M).
			MaxCost:     1 << 30, // maximum cost of cache (1GB).
			BufferItems: 64,      // number of keys per Get buffer.
		})
	return &RistrettoLocalCache{client: cache}
}
