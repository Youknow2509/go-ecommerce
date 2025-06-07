package initialize

import (
	"github.com/Youknow2509/go-ecommerce/internal/auth/controller/http"
	"github.com/gin-gonic/gin"
)

func initializeAuth(routerGroup *gin.RouterGroup) {
	// auth
	{
		http.AuthRouterManager.AuthPublicRouterGroup.InitAuthPublicRouter(routerGroup)
		http.AuthRouterManager.AuthPrivateRouterGroup.InitAuthPrivateRouter(routerGroup)
	}
}