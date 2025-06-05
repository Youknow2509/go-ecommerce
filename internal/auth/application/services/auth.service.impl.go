package services

import (
	"context"
	"github.com/Youknow2509/go-ecommerce/internal/auth/domain/cache"
	"github.com/Youknow2509/go-ecommerce/internal/auth/domain/models/entity"
	"github.com/Youknow2509/go-ecommerce/internal/auth/domain/responsitory"
)

// AuthService implements the IAuthentication interface
type AuthenticationService struct {
	responsitory     responsitory.IAuthResponsitory
	distributedCache cache.ICacheService
	localCache       cache.ICacheService
} 
// ##############################################################

// IsTwoFactorEnabled implements IAuthentication.
func (a AuthenticationService) IsTwoFactorEnabled(ctx context.Context, userID int) error {
	panic("unimplemented")
}

// Login implements IAuthentication.
func (a AuthenticationService) Login(ctx context.Context, input *entity.Login) (entity.Token, error) {
	panic("unimplemented")
}

// Logout implements IAuthentication.
func (a AuthenticationService) Logout(ctx context.Context, token *entity.Token) error {
	panic("unimplemented")
}

// RefreshToken implements IAuthentication.
func (a AuthenticationService) RefreshToken(ctx context.Context, token *entity.Token) (entity.Token, error) {
	panic("unimplemented")
}

// SetupTwoFactorAuth implements IAuthentication.
func (a AuthenticationService) SetupTwoFactorAuth(ctx context.Context, input *entity.SetupTwoFactorAuth) error {
	panic("unimplemented")
}

// VerifyTwoFactorAuth implements IAuthentication.
func (a AuthenticationService) VerifyTwoFactorAuth(ctx context.Context, input *entity.VerifyTwoFactorAuth) error {
	panic("unimplemented")
}

// ##############################################################

// new service and implement interface
func NewAuthService(
	responsitory responsitory.IAuthResponsitory,
	distributedCache cache.ICacheService,
	localCache cache.ICacheService,
) IAuthentication {
	return AuthenticationService{
		responsitory:     responsitory,
		distributedCache: distributedCache,
		localCache:       localCache,
	}
}
