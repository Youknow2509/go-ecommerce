package account

import (
	"log"

	"github.com/Youknow2509/go-ecommerce/internal/model"
	"github.com/Youknow2509/go-ecommerce/internal/service"
	"github.com/Youknow2509/go-ecommerce/internal/utils/context"
	"github.com/Youknow2509/go-ecommerce/response"
	"github.com/gin-gonic/gin"
)

var TwoFA = new(cUser2FA)

type cUser2FA struct {

}

// Setup 2fa godoc
// @Summary      Setup 2fa for user
// @Description  Setup 2fa for user
// @Tags         accounts 2fa		
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Authorization Bearer token"
// @Param        payload body model.SetupTwoFactorAuthInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrResponseData
// @Router       /v1/user/two-factor/setup [post]
func (a *cUser2FA) SetupTwoFactorAuth(ctx *gin.Context) {
	var params model.SetupTwoFactorAuthInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid params",
		})
		return
	}
	// get userId from uuid from token
	userId, err := context.GetUserIdFromUUID(ctx.Request.Context())
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	log.Println("userId:: ", userId)
	// add userId to params
	params.UserId = uint32(userId)
	// handle to service
	codeResult, err := service.UserLogin().SetupTwoFactorAuth(ctx, &params)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	response.SuccessResponse(ctx, codeResult, nil)
}

// @Summary      Verify user 2fa
// @Description  Verify user 2fa
// @Tags         accounts 2fa		
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Authorization Bearer token"
// @Param        payload body model.TwoFactorVerificationInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrResponseData
// @Router       /v1/user/two-factor/verify [post]
func (a *cUser2FA) VerifyTwoFactoryAuthentication(ctx *gin.Context) {
	var params model.TwoFactorVerificationInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid params",
		})
		return
	}
	// get userId from uuid from token
	userId, err := context.GetUserIdFromUUID(ctx.Request.Context())
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	log.Println("userId:: ", userId)
	// add userId to params
	params.UserId = uint32(userId)
	// handle to service
	codeResult, err := service.UserLogin().VerifyTwoFactorAuth(ctx, &params)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	response.SuccessResponse(ctx, codeResult, nil)
}