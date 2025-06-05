package responsitory

import (
	"context"

	"github.com/Youknow2509/go-ecommerce/internal/auth/domain/models/entity"
	rp "github.com/Youknow2509/go-ecommerce/internal/auth/domain/responsitory"
	"github.com/Youknow2509/go-ecommerce/internal/auth/infrastrutore/database"
)

// struct AuthResponsitory
type AuthResponsitory struct {
	db *database.Queries
}

// ####################################################

// BlockToken implements responsitory.IAuthResponsitory.
func (a *AuthResponsitory) BlockToken(ctx context.Context, accessToken string, refreshToken string) error {
	return a.db.BlockRefreshToken(
		ctx,
		database.BlockRefreshTokenParams{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	)
}

// BlockTokenWithId implements responsitory.IAuthResponsitory.
func (a *AuthResponsitory) BlockTokenWithId(ctx context.Context, tokenId int64) error {
	return a.db.BlockRefreshTokenWithId(
		ctx,
		tokenId,
	)
}

// CreateToken implements responsitory.IAuthResponsitory.
func (a *AuthResponsitory) CreateToken(ctx context.Context, input *entity.InputTokenCreate) error {
	return a.db.CreateToken(
		ctx,
		database.CreateTokenParams{
			AccessToken:  input.AccessToken,
			RefreshToken: input.RefreshToken,
		},
	)
}

// InfoToken implements responsitory.IAuthResponsitory.
func (a *AuthResponsitory) InfoToken(ctx context.Context, accessToken string, refreshToken string) (*entity.InputTokenCreate, error) {
	response, err := a.db.GetToken(
		ctx,
		database.GetTokenParams{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	)
	if err != nil {
		return nil, err
	}
	return &entity.InputTokenCreate{
		UserId:                uint32(response.UserID),
		AccessToken:           response.AccessToken,
		RefreshToken:          response.RefreshToken,
		AccessTokenExpiresAt:  response.AccessTokenExpiresAt.Unix(),
		RefreshTokenExpiresAt: response.RefreshTokenExpiresAt.Unix(),
	}, nil
}

// IsTokenRefreshedWithId implements responsitory.IAuthResponsitory.
func (a *AuthResponsitory) IsTokenRefreshedWithId(ctx context.Context, tokenId uint32) (bool, error) {
	response, err := a.db.IsTokenRefreshWithId(
		ctx,
		int64(tokenId),
	)
	if err != nil {
		return false, err
	}
	return response != 0, nil
}

// IsTokenRefreshedWithRefreshToken implements responsitory.IAuthResponsitory.
func (a *AuthResponsitory) IsTokenRefreshedWithRefreshToken(ctx context.Context, accessToken, refreshToken string) (bool, error) {
	response, err := a.db.IsTokenRefresh(
		ctx,
		database.IsTokenRefreshParams{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	)
	if err != nil {
		return false, err
	}
	return response != 0, nil
}

// ####################################################

// new and implement IAuthResponsitory
func NewAuthResponsitory(db *database.Queries) rp.IAuthResponsitory {
	return &AuthResponsitory{
		db: db,
	}
}
