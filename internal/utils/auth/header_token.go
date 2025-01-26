package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// Extract the token from the header
func ExtractBearerToken(c *gin.Context) (string, bool) {
	// Authorization: Bearer token
	authHeader := c.GetHeader("Authorization")
	if strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer "), true
	}
	return "", false
}