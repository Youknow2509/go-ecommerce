package account

import (

	"github.com/Youknow2509/go-ecommerce/internal/service"
	"github.com/Youknow2509/go-ecommerce/response"
	"github.com/gin-gonic/gin"
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