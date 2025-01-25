package account

import (
	"github.com/Youknow2509/go-ecommerce/internal/model"
	"github.com/Youknow2509/go-ecommerce/internal/service"
	"github.com/Youknow2509/go-ecommerce/response"
	"github.com/gin-gonic/gin"
)

var TwoFA = new(cUser2FA)

type cUser2FA struct {

}

func (a *cUser2FA) SetupTwoFactorAuth(ctx *gin.Context) {
	var params model.SetupTwoFactorAuthInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid params",
		})
		return
	}
	// get userId from uuid from token

	codeResult, err := service.UserLogin().SetupTwoFactorAuth(ctx, &params)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	response.SuccessResponse(ctx, codeResult, nil)
}