package services

import (
	"context"
	"errors"

	"github.com/Youknow2509/go-ecommerce/internal/auth/domain/models/entity"
)

// struct domain auth
type (
	// Authentication
	IAuthentication interface {
		Login(ctx context.Context, input *entity.Login) (entity.Token, error)
		Logout(ctx context.Context, token *entity.Token) error
		RefreshToken(ctx context.Context, token *entity.Token) (entity.Token, error)
		// v.v
		IsTwoFactorEnabled(ctx context.Context, userID int) error
		SetupTwoFactorAuth(ctx context.Context, input *entity.SetupTwoFactorAuth) error
		VerifyTwoFactorAuth(ctx context.Context, input *entity.VerifyTwoFactorAuth) error
	}

	// ##############################
	// Authorization
	IAuthorization interface {
		// v.v
	}

	// ##############################
	//
	// v.v.
)

// ##############################
// save variable interface helper create service
var (
	vIAuthentication IAuthentication
	vIAuthorization  IAuthorization
)

// ##############################
// initialize instance
func InitAuthenticationService(iAuthentication IAuthentication) {
	vIAuthentication = iAuthentication
}

// GetAuthenticationService return instance
func GetAuthenticationService() (IAuthentication, error) {
	if vIAuthentication == nil {
		return nil, errors.New("authentication service is not initialized")
	}
	return vIAuthentication, nil
}

// ##############################
// initialize instance
func InitAuthorizationService(iAuthorization IAuthorization) {
	vIAuthorization = iAuthorization
}

// GetAuthorizationService return instance
func GetAuthorizationService() (IAuthorization, error) {
	if vIAuthorization == nil {
		return nil, errors.New("authorization service is not initialized")
	}
	return vIAuthorization, nil
}
