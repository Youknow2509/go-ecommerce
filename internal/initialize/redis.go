package initialize

import (
	"context"
	"fmt"
	"strconv"

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