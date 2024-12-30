package user

import (
	// "github.com/Youknow2509/go-ecommerce/internal/controller"
	// "github.com/Youknow2509/go-ecommerce/internal/repo"
	// "github.com/Youknow2509/go-ecommerce/internal/service"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router

	// handle with non DJ
	// u := repo.NewUserRepository()
	// us := service.NewUserService(u)
	// userHanlderNonDependency := controller.NewUserController(us)

	// userRouterPublic := Router.Group("/user")
	// {
	// 	userRouterPublic.POST("/register", userHanlderNonDependency.Register)
	// 	// userRouterPublic.POST("/send_otp")
	// }

	// -> handle with wire go


	// private router
	userRouterPrivate := Router.Group("/user")
	// userRouterPrivate.Use(Limmited())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/get_info")
	}
}
