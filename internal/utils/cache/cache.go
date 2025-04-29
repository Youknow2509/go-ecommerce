package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/Youknow2509/go-ecommerce/internal/service"
	"github.com/redis/go-redis/v9"
)

/**
 * Get data in cache with key
 * Save data to obj use pointer
 */
func GetCache(ctx context.Context, key string, obj interface{}) error {
	rs, err := global.Rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return fmt.Errorf("cache not found: %s", key)
	} else if err != nil {
		return err
    }
	// convert rs to obj
	err = json.Unmarshal([]byte(rs), obj)
	if err != nil {
		return fmt.Errorf("unmarshal cache failed: %w", err)
    }
	return nil
}

/**
 * Set cache - set distributed cache and local cache
 * @param ctx
 * @param key
 * @param data
 * @param ttl
 * @param distributedCache
 * @param localCache
 * @return bool, error
 */
func SetCache(
	ctx context.Context, 
	key string, 
	data interface{}, 
	ttl int64,
	distributedCache service.IRedisCache,
	localCache service.ILocalCache,
) (res string, ok bool) {
	// timeTtl := time.Duration(ttl) * time.Second
	jsonData, _ := json.Marshal(data)
	if distributedCache != nil {
		err := distributedCache.Set(
			ctx, 
			key, 
			string(jsonData), 
			int(ttl),
		)
		if err != nil {
			return fmt.Sprintf("error setting distributed cache: %v", err), false
		}
	}

	if localCache != nil {
		okk := localCache.SetWithTTL(
			ctx, 
			key, 
			string(jsonData),
		)
		if !okk {
			return fmt.Sprintf("error setting local cache for key: %s", key), false
		}
	}

	return "", true
}