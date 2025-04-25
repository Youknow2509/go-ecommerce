package initialize

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

// Initial redis
func InitRedis() {
	r := global.Config.Redis
	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     r.Host + ":" + strconv.Itoa(r.Port),
		Password: r.Password, // no password set
		DB:       r.Database, // use default DB
		PoolSize: r.PoolSize, // default pool size - 10 connections in the pool
	})

	// Check connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Failed to connect to Redis, error: ", zap.Error(err))
		panic(err)
	}

	global.Rdb = rdb

	// redisExample()
}

// init redis sentinel
func InitRedisSentinel() {
	// connect to Redis Sentinel
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    global.Config.Redis.MasterName,
		SentinelAddrs: global.Config.Redis.SentinelAddrs,
		Password:      global.Config.Redis.Password,
		DB:            global.Config.Redis.Database,
		PoolSize:      global.Config.Redis.PoolSize,
		Dialer: func(ctx context.Context, network, addr string) (net.Conn, error) {
			// fmt.Println("Dialing to Redis Sentinel: ", addr)
			if global.Config.Server.Mode == "dev" {
				// Thay thế địa chỉ IP redis master bằng localhost (thay vì dùng ip trong docker)
				if strings.Contains(addr, strconv.Itoa(global.Config.Redis.Port)) {
					addr = global.Config.Redis.Host + ":" + strconv.Itoa(global.Config.Redis.Port)
				}
			}
			dialer := &net.Dialer{}
			return dialer.DialContext(ctx, network, addr)
		},
	})
	// Check connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Failed to connect to Redis Sentinel, error: ", zap.Error(err))
		panic(err)
	}
	fmt.Println("Connected to Redis Sentinel")

	// try setting and getting a value
	err = rdb.Set(ctx, "test_key", "Redis Sentinel", 0).Err()
	if err != nil {
		global.Logger.Error("Failed to set value in Redis Sentinel, error: ", zap.Error(err))
		panic(err)
	}
	value, err := rdb.Get(ctx, "test_key").Result()
	if err != nil {
		global.Logger.Error("Failed to get value from Redis Sentinel, error: ", zap.Error(err))
		panic(err)
	}
	fmt.Println("Value from Redis Sentinel: ", value)
	global.Rdb = rdb
	fmt.Println("Redis Sentinel initialized successfully")
	// Test redis sentinel
	// redisExample()
}

// Test redis
func redisExample() {
	err := global.Rdb.Set(ctx, "score", 100, 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("Set key 'score' to value '100'")

	value, err := global.Rdb.Get(ctx, "score").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("Get key 'score' value: ", value)
}
