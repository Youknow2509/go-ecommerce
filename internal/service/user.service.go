package service

import (
	rp "github.com/Bot-SomeOne/go-ecommerce/internal/repo"
)

type UserService struct {
	userRepo *rp.UserRepo
}

// NewUserService creates a new UserService
func NewUserService() *UserService {
	return &UserService{
		userRepo: rp.NewUserRepo(),
	}
}

// Get User Information Services
func (u *UserService) GetUserInfoService() string {
	return u.userRepo.GetInfoUser()
}