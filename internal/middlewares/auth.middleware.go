package middlewares

import (
	"context"
	"log"

	"github.com/Youknow2509/go-ecommerce/internal/utils/auth"
	"github.com/gin-gonic/gin"
)

// func authen middleware
func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get the url in request
		url := c.Request.URL.Path
		log.Printf("Request URL: %s", url)
		// get headers authorization - get jwt token in header
		jwtToken, err := auth.ExtractBearerToken(c)
		if !err {
			c.AbortWithStatusJSON(401, gin.H{"code": 40001, "err": "Unauthorized", "description": "Get authorization header failed"})
			return
		}
		// validate token
		claims, ok := auth.ValidateTokenSubject(jwtToken)
		if ok != nil {
			c.AbortWithStatusJSON(401, gin.H{"code": 40002, "err": "Invalid token", "description": "Validate token failed"})
			return
		}
		// update claims to context
		log.Println("Claims:: uuid:: ", claims.Subject)
		ctx := context.WithValue(c.Request.Context(), "SUBJECT_UUID", claims.Subject)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
