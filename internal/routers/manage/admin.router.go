package manage

import (
	"github.com/Youknow2509/go-ecommerce/internal/middlewares"
	"github.com/gin-gonic/gin"
)

type AdminRouter struct {
}

func (ur *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// public router
	adminRouterPublic := Router.Group("/admin")
	adminRouterPublic.Use(middlewares.PrometheusMiddleware())
	{
		adminRouterPublic.POST("/login")
	}

	// private router
	adminRouterPrivate := Router.Group("/admin/user")
	// adminRouterPrivate.Use(Limmited())
	// adminRouterPrivate.Use(Authen())
	// adminRouterPrivate.Use(Permission())
	{
		adminRouterPrivate.POST("/active_user")
	}
}
