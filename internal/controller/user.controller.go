package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Bot-SomeOne/go-ecommerce/internal/service"
)

// UserController struct
type UserController struct {
	userService *service.UserService
}

// NewUserController function
func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// controlelr -> service -> repo -> models -> dbs
func (u *UserController) GetUserByID(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": u.userService.GetUserInfoService(),
	})
}
