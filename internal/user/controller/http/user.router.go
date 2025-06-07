package http

import (
	"github.com/gin-gonic/gin"
)

// User Manager Router
type (
	// #############################

	UserPublicRouterGroup struct {
		handlerHttpController IUserHandlerHttp
	}

	// #############################

	UserPrivateRouterGroup struct {
		handlerHttpController IUserHandlerHttp	
	}

	// #############################
	UserRouterGroup struct {
		UserPublicRouterGroup
		UserPrivateRouterGroup
	}
)

var (
	UserRouterManager = &UserRouterGroup{}
	UserPublicRouter  = &UserPublicRouterGroup{}
	UserPrivateRouter = &UserPrivateRouterGroup{}
)

// #############################

// Init User Public Router
func (r *UserPublicRouterGroup) InitUserPublicRouter(router *gin.RouterGroup) {
	UserPublic := router.Group("/user")
	{
		UserPublic.POST("/register", r.handlerHttpController.RegisterUserController) // Register a new user
	}
}

// #############################

// Init User Private Router
func (r *UserPrivateRouterGroup) InitUserPrivateRouter(router *gin.RouterGroup) {
	UserPrivate := router.Group("/user")
	{
		UserPrivate.POST("/verify-register", r.handlerHttpController.VerifyRegisterUserController)          // Verify user registration
		UserPrivate.POST("/register-create-password", r.handlerHttpController.CreatePasswordUserController) // Create password for user after registration
	}
}

// #############################
// Init User Router Manager
func InitUserRouter(handlerHttpController IUserHandlerHttp) *UserRouterGroup {
	return &UserRouterGroup{
		UserPublicRouterGroup: UserPublicRouterGroup{
			handlerHttpController: handlerHttpController,
		},
		UserPrivateRouterGroup: UserPrivateRouterGroup{
			handlerHttpController: handlerHttpController,
		},
	}
}
