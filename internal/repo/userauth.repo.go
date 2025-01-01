package repo

import (
	"fmt"
	"time"

	"github.com/Youknow2509/go-ecommerce/global"
)

type IUserAuthRepository interface {
	AddOtp(email string, otp int, expirationTime int64) error
}

type userAuthRepository struct {
}

// new userAuthRepository creates a new userAuthRepository
func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}
}

// AddOtp implements IUserAuthRepository.
func (u *userAuthRepository) AddOtp(email string, otp int, expirationTime int64) error {
	key := fmt.Sprintf("usr:%s:otp", email) // usr:email:otp
	return global.Rdb.SetEx(ctx, key, otp, time.Duration(expirationTime)).Err()
}
