package http

import "github.com/gin-gonic/gin"

// interface handler http
type (
	IAuthHandlerHttp interface {
		LoginController(ctx *gin.Context)
		LogoutController(ctx *gin.Context)
		RefreshTokenController(ctx *gin.Context)
		//
		IsTwoFactorEnabledController(ctx *gin.Context)
		SetupTwoFactorAuthController(ctx *gin.Context)
		VerifyTwoFactorAuthController(ctx *gin.Context)
		// v.v
	}
	// ###############################
	AuthHandlerHttp struct{}
)
// ###############################

// IsTwoFactorEnabledController implements IAuthHandlerHttp.
func (a *AuthHandlerHttp) IsTwoFactorEnabledController(ctx *gin.Context) {
	panic("unimplemented")
}

// LoginController implements IAuthHandlerHttp.
func (a *AuthHandlerHttp) LoginController(ctx *gin.Context) {
	panic("unimplemented")
}

// LogoutController implements IAuthHandlerHttp.
func (a *AuthHandlerHttp) LogoutController(ctx *gin.Context) {
	panic("unimplemented")
}

// RefreshTokenController implements IAuthHandlerHttp.
func (a *AuthHandlerHttp) RefreshTokenController(ctx *gin.Context) {
	panic("unimplemented")
}

// SetupTwoFactorAuthController implements IAuthHandlerHttp.
func (a *AuthHandlerHttp) SetupTwoFactorAuthController(ctx *gin.Context) {
	panic("unimplemented")
}

// VerifyTwoFactorAuthController implements IAuthHandlerHttp.
func (a *AuthHandlerHttp) VerifyTwoFactorAuthController(ctx *gin.Context) {
	panic("unimplemented")
}

// ###############################

var (
	AuthHandlerHttpManager IAuthHandlerHttp = &AuthHandlerHttp{}
)

// ###############################
