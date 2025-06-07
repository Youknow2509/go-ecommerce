package initialize

import (
	"context"
	"fmt"
	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/redis/go-redis/v9"
	// "go.uber.org/zap"
)

var ctx = context.Background()

// init redis sentinel
func InitRedisSentinel() {
	// connect to Redis Sentinel
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    global.Config.Redis.MasterName,
		SentinelAddrs: global.Config.Redis.SentinelAddrs,
		Password:      global.Config.Redis.Password,
		DB:            global.Config.Redis.Database,
		PoolSize:      global.Config.Redis.PoolSize,
	})
	// Check connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		// global.Logger.Error("Failed to connect to Redis Sentinel, error: ", zap.Error(err))
		panic(err)
	}
	fmt.Println("Connected to Redis Sentinel")

	// try setting and getting a value
	err = rdb.Set(ctx, "test_key", "Redis Sentinel", 0).Err()
	if err != nil {
		// global.Logger.Error("Failed to set value in Redis Sentinel, error: ", zap.Error(err))
		panic(err)
	}
	value, err := rdb.Get(ctx, "test_key").Result()
	if err != nil {
		// global.Logger.Error("Failed to get value from Redis Sentinel, error: ", zap.Error(err))
		panic(err)
	}
	fmt.Println("Value from Redis Sentinel: ", value)
	global.Rdb = rdb
	fmt.Println("Redis Sentinel initialized successfully")
}