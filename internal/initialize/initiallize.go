package initialize

import (
	"github.com/Youknow2509/go-ecommerce/internal/auth/controller/http"
	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine {
	
	// Initialize the Gin router
	router := gin.Default()

	// Set up routes
	manager := router.Group("/api/v1")
	{
		// auth
		http.AuthRouterManager.AuthPublicRouterGroup.InitAuthPublicRouter(manager)
		http.AuthRouterManager.AuthPrivateRouterGroup.InitAuthPrivateRouter(manager)
	}

	return router
}