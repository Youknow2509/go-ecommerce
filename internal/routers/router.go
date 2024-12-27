package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/ping", PongHandler) // /v1/ping
		v1.PUT("/ping", PongHandler)
		v1.POST("/ping", PongHandler)
		v1.DELETE("/ping", PongHandler)
		v1.OPTIONS("/ping", PongHandler)
		v1.PATCH("/ping", PongHandler)
		v1.HEAD("/ping", PongHandler)
	}

	router.NoRoute(PageNotFoundHandler)

	return router
}

func PongHandler(c *gin.Context) {
	name := c.Param("name")
	uid := c.Query("uid")
	c.JSON(http.StatusNotFound, gin.H{
		"message": "pong",
		"name":    name,
		"uid":     uid,
	})
}

func PageNotFoundHandler(c *gin.Context) {
	name := c.Param("name")
	uid := c.Query("uid")
	c.JSON(http.StatusNotFound, gin.H{
		"message": "404 page not found",
		"name":    name,
		"uid":     uid,
	})
}
