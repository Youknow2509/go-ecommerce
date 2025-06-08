package initialize

import (
	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine {
	InitializeConfig()
	InitializeLogger()
	InitRedisSentinel()
	InitializeMysqlAll()
	InitKafka()
	// Initialize the Gin router
	router := gin.Default()
	// Set up routes
	manager := router.Group("/api/v1")

	// initializeUser(manager)
	InitializeUser(manager)

	return router
}
