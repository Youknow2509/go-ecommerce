package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Youknow2509/go-ecommerce/global"
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