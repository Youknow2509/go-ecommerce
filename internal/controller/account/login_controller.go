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