package repo

import (
	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/Youknow2509/go-ecommerce/internal/model"
)

// Interface Version
type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {

}

func NewUserRepository() IUserRepository {
    return &userRepository{}
}

func (u *userRepository) GetUserByEmail(email string) bool {
	// SELECT * FROM users WHERE email = '??' ORDER BY email
	row := global.Mdb.Table(TABLE_NAME_GOCRMUSER).Where("usr_email = ?", email).First(&model.GoCrmUser{}).RowsAffected
	return row != NUMBER_NULL
}
