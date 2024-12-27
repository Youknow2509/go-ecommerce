package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct {
}

// NewPongController creates a new PongController
func NewPongController() *PongController {
	return &PongController{}
}

func (p *PongController) PongHandler(c *gin.Context) {
	fmt.Println("--> My Handler PongController")
	name := c.Param("name")
	uid := c.Query("uid")
	c.JSON(http.StatusNotFound, gin.H{
		"message": "pong",
		"name":    name,
		"uid":     uid,
	})
}
