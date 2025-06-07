package services

import (
	"context"
	"fmt"

	"github.com/Youknow2509/go-ecommerce/internal/user/application/model"
)

// define interface service
type (
	IUserService interface {
		RegisterUserService(ctx context.Context, input *model.InputUserRegister) (int, error)
		VerifyRegisterUserService(ctx context.Context, input *model.InputUserVerifyRegister) (*model.OutputUserVerifyRegister, int, error)
		CreatePasswordUserService(ctx context.Context, input *model.InputUserCreatePassword) (int, error)
	}
)
// ##################################################################

var (
	vIUserService IUserService
)
// ##################################################################

// Init IUserService
func InitUserService(i IUserService) {
	vIUserService = i
}

// Get IUserService
func GetUserService() (IUserService, error) {
	if vIUserService == nil {
		return nil, fmt.Errorf("UserService is not initialized, please call InitUserService first")
	}
	return vIUserService, nil
}