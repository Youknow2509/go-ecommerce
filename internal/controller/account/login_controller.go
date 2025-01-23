package account

import (
	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/Youknow2509/go-ecommerce/internal/model"
	"github.com/Youknow2509/go-ecommerce/internal/service"
	"github.com/Youknow2509/go-ecommerce/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Login = new(cUserLogin)

type cUserLogin struct {

}

func (cU *cUserLogin) Login(c *gin.Context) {

	err := service.UserLogin().Login(c)
	if err != nil {
        response.ErrorResponse(c, response.ErrCodeParamInvalid, err.Error())
        return
    }

	response.SuccessResponse(c, response.ErrCodeSuccess, nil)
}

// Register godoc
// @Summary      Register a new account
// @Description  When user register, system will send OTP to user's phone number or email address
// @Tags         accounts management
// @Accept       json
// @Produce      json
// @Param        payload body model.RegisterInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrResponseData
// @Router       /v1/user/register [post]
func (cU *cUserLogin) Register(c *gin.Context) {
	var params model.RegisterInput
	if err := c.ShouldBindJSON(&params); err!= nil {
        response.ErrorResponse(c, response.ErrCodeParamInvalid, err.Error())
        return
    }

	codeStatus, err := service.UserLogin().Register(c, &params)
	if err != nil {
		global.Logger.Error("Error registering user otp", zap.Error(err))
		response.ErrorResponse(c, codeStatus, err.Error())
		return
	}

	response.SuccessResponse(c, response.ErrCodeSuccess, nil)
}

// Verify login by user 
// @Summary      Verify OTP bu user when register	
// @Description  Verify OTP bu user when register
// @Tags         accounts management
// @Accept       json
// @Produce      json
// @Param        payload body model.RegisterInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrResponseData
// @Router       /v1/user/verify_account [post]
func (cU *cUserLogin) VerifyOTP(c *gin.Context) {
	var params model.VerifyInput
	if err := c.ShouldBindJSON(&params); err!= nil {
		response.ErrorResponse(c, response.ErrCodeParamInvalid, err.Error())
		return
	}

	out, err := service.UserLogin().VerifyOTP(c, &params)
	if err != nil {
		global.Logger.Error("Error verifying user otp", zap.Error(err))
		response.ErrorResponse(c, response.ErrCodeParamInvalid, err.Error())
		return
	}

	response.SuccessResponse(c, response.ErrCodeSuccess, out)
}