package http

import (
	"github.com/Youknow2509/go-ecommerce/internal/user/application/model"
	"github.com/Youknow2509/go-ecommerce/internal/user/application/services"
	"github.com/Youknow2509/go-ecommerce/internal/user/controller/dto"
	"github.com/Youknow2509/go-ecommerce/response"
	"github.com/gin-gonic/gin"
)

// interface handler http
type (
	IUserHandlerHttp interface {
		RegisterUserController(c *gin.Context)       // Register a new user
		VerifyRegisterUserController(c *gin.Context) // Verify user registration
		CreatePasswordUserController(c *gin.Context) // Create password for user after registration
		// v.v
	}
	// ###############################
	UserHandlerHttp struct {
		service services.IUserService
	}
)

// ###############################

// CreatePasswordUserController implements IUserHandlerHttp.
//
//				@Description Create password for user after registration
//				@Tags UserManagement
//				@Accept json
//				@Produce json
//				@Param request body dto.UserCreatePasswordDto true "User create password data"
//		     	@Success		200		{object}	response.ResponseData "Success response"
//	         @Failure		400		{object}	response.ErrResponseData "Bad request"
//	         @Failure		500		{object}	response.ErrResponseData "Internal server error"
//	         @Failure		404		{object}	response.ErrResponseData "Not found"
//				@Router /api/v1/user/register-create-password [post]
func (u *UserHandlerHttp) CreatePasswordUserController(c *gin.Context) {
	// bind request
	request := &dto.UserCreatePasswordDto{}
	if err := c.ShouldBindJSON(request); err != nil {
		response.ErrorResponse(c, response.ErrCodeBindRegisterInput, err.Error())
		return
	}
	// call user service
	codeReponse, err := u.service.CreatePasswordUserService(
		c,
		&model.InputUserCreatePassword{
			AccountName:         request.AccountName,
			TokenCreatePassword: request.TokenCreatePassword,
			Password:            request.Password,
		},
	)
	if err != nil {
		response.ErrorResponse(c, codeReponse, err.Error())
		return
	}
	// response success
	response.SuccessResponse(c, codeReponse, "User password created successfully")
}

// RegisterUserController implements IUserHandlerHttp.
//
//				@Description Register a new user
//				@Tags UserManagement
//				@Accept json
//				@Produce json
//				@Param request body dto.UserRegisterDto true "User registration data"
//		     	@Success		200		{object}	response.ResponseData "Success response"
//	         @Failure		400		{object}	response.ErrResponseData "Bad request"
//	         @Failure		500		{object}	response.ErrResponseData "Internal server error"
//	         @Failure		404		{object}	response.ErrResponseData "Not found"
//				@Router /api/v1/user/register [post]
func (u *UserHandlerHttp) RegisterUserController(c *gin.Context) {
	// bind request
	request := &dto.UserRegisterDto{}
	if err := c.ShouldBindJSON(request); err != nil {
		response.ErrorResponse(c, response.ErrCodeBindRegisterInput, err.Error())
		return
	}
	// call user service
	codeReponse, err := u.service.RegisterUserService(
		c,
		&model.InputUserRegister{
			AcountName: request.AcountName,
			VerifyType: request.VerifyType,
		},
	)
	if err != nil {
		response.ErrorResponse(c, codeReponse, err.Error())
		return
	}
	// response success
	response.SuccessResponse(c, codeReponse, "User registered successfully")
}

// VerifyRegisterUserController implements IUserHandlerHttp.
//
//				@Description Verify user registration
//				@Tags UserManagement
//				@Accept json
//				@Produce json
//				@Param request body dto.UserVerifyRegisterDto true "User registration verification data"
//		     	@Success		200		{object}	response.ResponseData "Success response"
//	         @Failure		400		{object}	response.ErrResponseData "Bad request"
//	         @Failure		500		{object}	response.ErrResponseData "Internal server error"
//	         @Failure		404		{object}	response.ErrResponseData "Not found"
//				@Router /api/v1/user/verify-register [post]
func (u *UserHandlerHttp) VerifyRegisterUserController(c *gin.Context) {
	// bind request
	request := &dto.UserVerifyRegisterDto{}
	if err := c.ShouldBindJSON(request); err != nil {
		response.ErrorResponse(c, response.ErrCodeBindRegisterInput, err.Error())
		return
	}
	// call user service
	responseData, codeReponse, err := u.service.VerifyRegisterUserService(
		c,
		&model.InputUserVerifyRegister{
			AcountName: request.AcountName,
			VerifyCode: request.VerifyCode,
		},
	)
	if err != nil {
		response.ErrorResponse(c, codeReponse, err.Error())
		return
	}
	// response success
	response.SuccessResponse(c, codeReponse, responseData)
}

// ###############################

var (
	UserHandlerHttpManager IUserHandlerHttp
)

// ###############################

func NewUserHandlerHttp(userService services.IUserService) IUserHandlerHttp {
	return &UserHandlerHttp{
		service: userService,
	}
}
