package middlewares

import (
	"fmt"
	"log"
	"time"

	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/Youknow2509/go-ecommerce/internal/utils/context"
	"github.com/Youknow2509/go-ecommerce/response"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	redisStore "github.com/ulule/limiter/v3/drivers/store/redis"
)

// struct Rate Limiter
type RateLimiter struct {
	globalRateLimiter         *limiter.Limiter
	publicAPIRateLimiter      *limiter.Limiter
	userPrivateAPIRateLimiter *limiter.Limiter
}

// new RateLimiter
func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		globalRateLimiter:         rateLimiter("100-S"),
		publicAPIRateLimiter:      rateLimiter("80-S"),
		userPrivateAPIRateLimiter: rateLimiter("50-S"),
	}
}

/**
 * create limiter with time interval
 * Example:
 * 	5 reqs/second: "5-S"
 * 	10 reqs/minute: "10-M"
 * 	1000 reqs/hour: "1000-H"
 * 	2000 reqs/day: "2000-D"
 */
func rateLimiter(interval string) *limiter.Limiter {
	store, err := redisStore.NewStoreWithOptions(
		global.Rdb,
		limiter.StoreOptions{
			Prefix:          "rate-limiter", // prefix every key with "rate-limiter"
			MaxRetry:        3,              // maximum number of
			CleanUpInterval: time.Hour,      // clean up expired data
		},
	)
	if err != nil {
		return nil
	}
	rate, err := limiter.NewRateFromFormatted(interval)
	if err != nil {
		return nil
	}
	instace := limiter.New(store, rate)
	return instace
}

/**
 * global limiter
 */
func (rl *RateLimiter) GlobalLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := "global"
		log.Printf("global limiter")
		limiterContext, err := rl.globalRateLimiter.Get(c, key)
		if err != nil {
			c.JSON(response.ErrCodeTooManyRequests, gin.H{"error": "Failed to get global rate limiter"})
			c.Next()
			return
		}
		if limiterContext.Reached {
			c.JSON(response.ErrCodeTooManyRequests, gin.H{"error": "Rate limit breached global, try again"})
			c.AbortWithStatusJSON(response.ErrCodeTooManyRequests, gin.H{"error": "Rate limit breached global, try again"})
			return
		}
		c.Next()
	}
}

/**
 * public api limiter
 */
func (rl *RateLimiter) PublicAPILimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		urlPath := c.Request.URL.Path
		rateLimitPath := rl.fillterLimitUrlPath(urlPath)
		if rateLimitPath != nil {
			log.Println("Client ip:: ", c.ClientIP())
			key := fmt.Sprintf("%s-%s", "111-222-333-44", urlPath)
			limiterContext, err := rateLimitPath.Get(c, key)
			if err != nil {
				fmt.Println("Failed to get public api rate limiter")
				c.Next()
				return
			}
			if limiterContext.Reached {
				c.JSON(response.ErrCodeTooManyRequests, gin.H{"error": "Rate limit breached public api, try again"})
				c.AbortWithStatusJSON(response.ErrCodeTooManyRequests, gin.H{"error": "Rate limit breached public api, try again"})
				return
			}
		}
	}
}

/**
 * user private api limiter
 */
func (rl *RateLimiter) UserPrivateAPILimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		urlPath := c.Request.URL.Path
		rateLimitPath := rl.fillterLimitUrlPath(urlPath)
		if rateLimitPath != nil {
			userId, err := context.GetUserIdFromUUID(c)
			if err != nil {
				c.JSON(response.ErrCodeAuthFailed, gin.H{"error": "Failed to get user id from token"})
				c.AbortWithStatusJSON(response.ErrCodeAuthFailed, gin.H{"error": "Failed to get user id from token"})
				c.Next()
				return
			}
			key := fmt.Sprintf("%d-%s", userId, urlPath)
			limiterContext, err := rateLimitPath.Get(c, key)
			if err != nil {
				fmt.Println("Failed to get private api rate limiter")
				c.Next()
				return
			}
			if limiterContext.Reached {
				c.JSON(response.ErrCodeTooManyRequests, gin.H{"error": "Rate limit breached private api, try again"})
				c.AbortWithStatusJSON(response.ErrCodeTooManyRequests, gin.H{"error": "Rate limit breached private api, try again"})
				return
			}
		}
	}
}

// fillterLimitUrlPath
func (rl *RateLimiter) fillterLimitUrlPath(urlPath string) *limiter.Limiter {
	if urlPath == "/metrics" {
		return nil
	} else if urlPath == "/v1/user/login"{
		return rl.publicAPIRateLimiter
	} else if urlPath == "/v1/user/info" {
		return rl.userPrivateAPIRateLimiter
	} else {
		return rl.globalRateLimiter
	}
}
