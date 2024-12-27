package routers

import (
	"github.com/gin-gonic/gin"
	c "github.com/Bot-SomeOne/go-ecommerce/internal/controller"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/ping", c.NewPongController().PongHandler) // /v1/ping
		v1.GET("/user", c.NewUserController().GetUserByID) // /v1/user
		// v1.PUT("/ping", controller.NewPongController().PongHandler)
		// v1.POST("/ping", controller.NewPongController().PongHandler)
		// v1.DELETE("/ping", controller.NewPongController().PongHandler)
		// v1.OPTIONS("/ping", controller.NewPongController().PongHandler)
		// v1.PATCH("/ping", controller.NewPongController().PongHandler)
		// v1.HEAD("/ping", controller.NewPongController().PongHandler)
	}

	return router
}
