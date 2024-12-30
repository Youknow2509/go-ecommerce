package controller

import (
	"github.com/Youknow2509/go-ecommerce/internal/service"
	"github.com/Youknow2509/go-ecommerce/response"
	"github.com/gin-gonic/gin"
)

// UserController struct
type UserController struct {
	userService service.IUserService
}

// NewUserController function
func NewUserController(
	userService service.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

// Register user
func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.RegisterService("","")
	response.SuccessResponse(c, result, "")
}

// controlelr -> service -> repo -> models -> dbs
func (u *UserController) GetUserByID(c *gin.Context) {
	// response.SuccessResponse(c, 20001, u.userService.GetUserInfoService())
	response.ErrorResponse(c, 20003, "")
}
