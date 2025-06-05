package http

import "github.com/gin-gonic/gin"

// Auth Manager Router
type (
	// #############################

	AuthPublicRouterGroup struct {
	}

	// #############################

	AuthPrivateRouterGroup struct {
	}

	// #############################
	AuthRouterGroup struct {
		AuthPublicRouterGroup
		AuthPrivateRouterGroup
	}
)

var (
	AuthRouterManager = &AuthRouterGroup{}
	AuthPublicRouter  = &AuthPublicRouterGroup{}
	AuthPrivateRouter = &AuthPrivateRouterGroup{}
)

// #############################

// Init Auth Public Router
func (r *AuthPublicRouterGroup) InitAuthPublicRouter(router *gin.RouterGroup) {
	authPublic := router.Group("/auth")
	{
		authPublic.POST("/login", AuthHandlerHttpManager.LoginController)
		authPublic.POST("/logout", AuthHandlerHttpManager.LogoutController)
	}
}

// #############################

// Init Auth Private Router
func (r *AuthPrivateRouterGroup) InitAuthPrivateRouter(router *gin.RouterGroup) {
	authPrivate := router.Group("/auth")
	{
		authPrivate.POST("/refresh-token", AuthHandlerHttpManager.RefreshTokenController)

		// Two Factor Auth
		authPrivate.GET("/2fa/is-enabled", AuthHandlerHttpManager.IsTwoFactorEnabledController)
		authPrivate.POST("/2fa/setup", AuthHandlerHttpManager.SetupTwoFactorAuthController)
		authPrivate.POST("/2fa/verify", AuthHandlerHttpManager.VerifyTwoFactorAuthController)
	}
}

// #############################
