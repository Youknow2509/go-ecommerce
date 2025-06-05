package responsitory

import (
	"context"
	"fmt"

	"github.com/Youknow2509/go-ecommerce/internal/auth/domain/models/entity"
)

// interface responsitory auth
type (
	IAuthResponsitory interface {
		CreateToken(ctx context.Context, input *entity.InputTokenCreate) error
		IsTokenRefreshedWithId(ctx context.Context, tokenId uint32) (bool, error)
		IsTokenRefreshedWithRefreshToken(ctx context.Context, accessToken, refreshToken string) (bool, error)
		//
		InfoToken(ctx context.Context, accessToken, refreshToken string) (*entity.InputTokenCreate, error)
		//
		BlockToken(ctx context.Context, accessToken, refreshToken string) error
		BlockTokenWithId(ctx context.Context, tokenId int64) error
		// v.v
	}
)

// ##################################################

// variable for service
var (
	vAuthResponsitory IAuthResponsitory
)

// ##################################################
// get instance of auth responsitory
func GetAuthResponsitory() (IAuthResponsitory, error) {
	if vAuthResponsitory == nil {
		return nil, fmt.Errorf("auth responsitory is not initialized")// TODO handle consts error
	}
	return vAuthResponsitory, nil
}

// initialize auth responsitory
func InitAuthResponsitory(responsitory IAuthResponsitory) {
	vAuthResponsitory = responsitory
}