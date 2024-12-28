package initialize

import (
	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/Youknow2509/go-ecommerce/internal/routers"
	"github.com/gin-gonic/gin"
)

// Initial router
func InitRouter() *gin.Engine {
	// router := gin.Default()
	var router *gin.Engine

	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		router = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
	}

	// middlewares
	router.Use() // logger
	router.Use() // cross
	router.Use() // limiter global

	manageRouter := routers.RouterGroupApp.Manage
	userRouter := routers.RouterGroupApp.User

	MainGroup := router.Group("/v1")
	{
		MainGroup.GET("/checkStatus") // tracking monitor 
	}
	{
		userRouter.InitUserRouter(MainGroup)
        userRouter.InitProductRouter(MainGroup)
        //... other routes...
	}
	{
		manageRouter.InitAdminRouter(MainGroup)
		manageRouter.InitUserRouter(MainGroup)
		//... other routes...
	}

	return router
}
