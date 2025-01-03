package repo

import (
	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/Youknow2509/go-ecommerce/internal/database"
)

// Interface Version
type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {
	sqlc *database.Queries
}

func NewUserRepository() IUserRepository {
    return &userRepository{
		sqlc: database.New(global.Mdbc),
	}
}

func (u *userRepository) GetUserByEmail(email string) bool {
	// SELECT * FROM users WHERE email = '??' ORDER BY email
	// row := global.Mdb.Table(TABLE_NAME_GOCRMUSER).Where("usr_email = ?", email).First(&model.GoCrmUser{}).RowsAffected
	row, err := u.sqlc.GetUserByEmailSQLC(ctx, email)
	if err != nil {
		return false
	}

	return row.UsrID != 0
}
