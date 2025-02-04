package controller

import (
	"fmt"

	"github.com/Youknow2509/go-ecommerce/internal/service"
	"github.com/Youknow2509/go-ecommerce/internal/vo"
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
// @Summary      Register a new account - Use microservice handle to kafka
// @Description  When user register, system will send OTP to user's phone number or email address
// @Tags         accounts management
// @Accept       json
// @Produce      json
// @Param        payload body vo.UserRegisterRequest true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrResponseData
// @Router       /v1/user/register [post]
func (uc *UserController) Register(c *gin.Context) {
	var params vo.UserRegisterRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(c, response.ErrCodeParamInvalid, err.Error())
		return
	}
	fmt.Printf("Email param:: %s\n", params.Email)
	result := uc.userService.RegisterService(params.Email, params.Purpose)
	
	response.SuccessResponse(c, result, nil)
}

// controlelr -> service -> repo -> models -> dbs
func (u *UserController) GetUserByID(c *gin.Context) {
	// response.SuccessResponse(c, 20001, u.userService.GetUserInfoService())
	response.ErrorResponse(c, 20003, "")
}
