package service

import "context"

// Interface for local cache
type ILocalCache interface {
	Get(ctx context.Context, key string) (interface{}, bool)
	Set(ctx context.Context, key string, value interface{}) bool
	SetWithTTL(ctx context.Context, key string, value interface{}) bool
	Del(ctx context.Context, key string) bool
}

// var localCache ILocalCache
var vLocalCache ILocalCache

// Init Singleton for local cache
func InitLocalCache(localCache ILocalCache) {
	if vLocalCache == nil {
		vLocalCache = localCache
	}
}

// GetLocalCache returns the local cache instance
func GetLocalCache() ILocalCache {
	if vLocalCache == nil {
		panic("local cache not initialized")
	}
	return vLocalCache
}
