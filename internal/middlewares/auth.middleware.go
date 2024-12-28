package middlewares

import (
	"github.com/Youknow2509/go-ecommerce/response"
	"github.com/gin-gonic/gin"
)

// func authen middleware
func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrInvalidToken, "")
			c.Abort()
			return
		}
		c.Next()
		// Something after next
	}
}
