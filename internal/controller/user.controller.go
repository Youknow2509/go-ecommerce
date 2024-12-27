package controller

import (
	"github.com/Bot-SomeOne/go-ecommerce/internal/service"
	"github.com/Bot-SomeOne/go-ecommerce/response"
	"github.com/gin-gonic/gin"
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
	// response.SuccessResponse(c, 20001, u.userService.GetUserInfoService())
	response.ErrorResponse(c, 20003, "")
}
