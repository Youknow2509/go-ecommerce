package impl

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Youknow2509/go-ecommerce/internal/consts"
	"github.com/Youknow2509/go-ecommerce/internal/service"
	"github.com/dgraph-io/ristretto/v2"
)

// restretto cache impl ILocalCache
type RestrettoCache struct {
	cache *ristretto.Cache[string, string]
}

// Del implements service.ILocalCache.
func (r *RestrettoCache) Del(ctx context.Context, key string) bool {
	_, found := r.cache.Get(key)
	if !found {
		return false
	}
	r.cache.Del(key)
	return true
}

// Get implements service.ILocalCache.
func (r *RestrettoCache) Get(ctx context.Context, key string) (interface{}, bool) {
	value, found := r.cache.Get(key)
	if !found {
		return nil, false
	}
	var data interface{}
	err := json.Unmarshal([]byte(value), &data)
	if err != nil {
		return nil, false
	}
	return data, true
}

// Set implements service.ILocalCache.
func (r *RestrettoCache) Set(ctx context.Context, key string, value interface{}) bool {
	jsonData, _ := json.Marshal(value)
	return r.cache.Set(key, string(jsonData), 1)
}

// SetWithTTL implements service.ILocalCache.
func (r *RestrettoCache) SetWithTTL(ctx context.Context, key string, value interface{}) bool {
	timeTTL := time.Minute * time.Duration(consts.TIME_TTL_LOCAL_CACHE)
	jsonData, _ := json.Marshal(value)
	return r.cache.SetWithTTL(key, string(jsonData), 1, timeTTL)
}

// new restretto cache
func NewRestrettoCache() service.ILocalCache {
	cache, err := ristretto.NewCache(
		&ristretto.Config[string, string]{
			NumCounters: 1e7,     // number of keys to track frequency of (10M).
			MaxCost:     1 << 30, // maximum cost of cache (1GB).
			BufferItems: 64,      // number of keys per Get buffer.
		})
	if err != nil {
		panic(err)
	}
	return &RestrettoCache{cache: cache}
}
