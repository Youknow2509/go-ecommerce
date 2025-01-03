package impl

import (
	"context"

	"github.com/Youknow2509/go-ecommerce/internal/database"
	"github.com/Youknow2509/go-ecommerce/internal/service"
)

// struct
type sUserLogin struct {
	r *database.Queries
}

// new sUserLogin implementation interface for IUserLogin
func NewSUserLogin(r *database.Queries) service.IUserLogin {
	return &sUserLogin{
		r: r,
	}
}

// Login implements service.IUserLogin.
func (s *sUserLogin) Login(ctx context.Context) error {
	return nil
}

// Register implements service.IUserLogin.
func (s *sUserLogin) Register(ctx context.Context) error {
	return nil
}

// VerifyOTP implements service.IUserLogin.
func (s *sUserLogin) VerifyOTP(ctx context.Context) error {
	return nil
}
