package user

import (
	// "github.com/Youknow2509/go-ecommerce/internal/controller"
	// "github.com/Youknow2509/go-ecommerce/internal/repo"
	// "github.com/Youknow2509/go-ecommerce/internal/service"
	"github.com/Youknow2509/go-ecommerce/internal/controller/account"
	"github.com/Youknow2509/go-ecommerce/internal/middlewares"
	// "github.com/Youknow2509/go-ecommerce/internal/wire"
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
	// userController, _ := wire.InitUserRouterHandle()
	userRouterPublic := Router.Group("/user")
	{
		// userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/register", account.Login.Register)
		userRouterPublic.POST("/login", account.Login.Login)
		userRouterPublic.POST("/verify_account", account.Login.VerifyOTP)	
		userRouterPublic.POST("/upgrade_password_register", account.Login.UpgradePasswordRegister)	
		
		// userRouterPublic.POST("/send_otp")
	}

	// private router
	userRouterPrivate := Router.Group("/user")
	userRouterPrivate.Use(middlewares.AuthenMiddleware())
	// userRouterPrivate.Use(Limmited())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/get_info")
		userRouterPrivate.POST("/two-factor/setup", account.TwoFA.SetupTwoFactorAuth)
		userRouterPrivate.POST("/two-factor/verify", account.TwoFA.VerifyTwoFactoryAuthentication)

	}
}
