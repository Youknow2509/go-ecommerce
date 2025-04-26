package service

import (
	"context"

	"github.com/Youknow2509/go-ecommerce/internal/model"
)

// create interface
type (
	IUserLogin interface {
		Login(ctx context.Context, in *model.LoginInput) (codeResult int, out model.LoginOutput, err error)
		Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error)
		VerifyOTP(ctx context.Context, in *model.VerifyInput) (out model.VerifyOTPOutput, err error)
		UpdatePasswordRegister(ctx context.Context, in *model.UpdatePasswordInput) (userId int, err error)
	
		// two-factor authentication
		IsTwoFactorEnabled(ctx context.Context, userId int) (codeResult int, err error)
		// setup authentication
		SetupTwoFactorAuth(ctx context.Context, in *model.SetupTwoFactorAuthInput) (codeResult int, err error)
		// verify authentication
		VerifyTwoFactorAuth(ctx context.Context, in *model.TwoFactorVerificationInput) (codeResult int, err error)
	
	}

	IUserInfo interface {
		GetInfoByUserId(ctx context.Context) error
		GetAllUser(ctx context.Context) error
	}

	IUserAdmin interface {
		RemoveUser(ctx context.Context) error
		FindOneUser(ctx context.Context) error
	}
)

// variables for service interface
var (
	localUserLogin IUserLogin
	localUserInfo IUserInfo
	localUserAdmin IUserAdmin
)

/**
 * Handle interface IUserLogin
 */
// Get interface IUser
func UserLogin() IUserLogin {
	if localUserLogin == nil {
        panic("implement localuserlogin not found for interface IUserLogin")
    }
    return localUserLogin
}

// Init interface IUserLogin
func InitUserLogin(userLogin IUserLogin) {
    localUserLogin = userLogin
}

/**
 * Handle interface IUserInfo
 */
// Get interface IUserInfo
func UserInfo() IUserInfo {
	if localUserInfo == nil {
        panic("implement localuserInfo not found for interface IUserInfo")
    }
    return localUserInfo
}

// Init interface IUserInfo
func InitUserInfo(userInfo IUserInfo) {
    localUserInfo = userInfo
}

/**
 * Handle interface IUserAdmin
 */
// Get interface IUserAdmin
func UserAdmin() IUserAdmin {
	if localUserAdmin == nil {
        panic("implement localuserAdmin not found for interface IUserAdmin")
    }
    return localUserAdmin
}

// Init interface IUserAdmin
func InitUserAdmin(userAdmin IUserAdmin) {
    localUserAdmin = userAdmin
}
