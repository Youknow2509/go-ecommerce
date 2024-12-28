package routers

import (
	"fmt"

	c "github.com/Youknow2509/go-ecommerce/internal/controller"
	"github.com/Youknow2509/go-ecommerce/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> AA")
		c.Next()
		fmt.Println("After --> AA")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> BB")
		c.Next()
		fmt.Println("After --> BB")
	}
}

func CC() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> CC")
		c.Next()
		fmt.Println("After --> CC")
	}
}

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.AuthenMiddleware(), AA(), BB(), CC())

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
