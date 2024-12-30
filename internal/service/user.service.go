package service

import (
	rp "github.com/Youknow2509/go-ecommerce/internal/repo"
	"github.com/Youknow2509/go-ecommerce/response"
)

// type UserService struct {
// 	userRepo *rp.UserRepo
// }

// // NewUserService creates a new UserService
// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: rp.NewUserRepo(),
// 	}
// }

// // Get User Information Services
// func (u *UserService) GetUserInfoService() string {
// 	return u.userRepo.GetInfoUser()
// }

// Interface Version
type IUserService interface {
	RegisterService(email string, purpose string) int
}

type userService struct {
	userRepo rp.IUserRepository
	// ...
}

func NewUserService(u rp.IUserRepository) IUserService {
	return &userService{
		userRepo: u,
	}
}

// RegisterService implements IUserService.
func (u *userService) RegisterService(email string, purpose string) int {
	// Check email exists
	if u.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExist
	}
	
	return response.ErrCodeSuccess
}
