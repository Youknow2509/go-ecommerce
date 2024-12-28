package manage

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	// userRouterPublic := Router.Group("/admin/user")
	// {
	// 	userRouterPublic.POST("/register")
	// 	userRouterPublic.POST("/send_otp")
	// }

	// private router
	// userRouterPrivate := Router.Group("/admin/user")
	// userRouterPrivate.Use(Limmited())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	// {
	// 	userRouterPrivate.POST("/active_user")
	// }
}
