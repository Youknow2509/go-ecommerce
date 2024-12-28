package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"column:uuid; type:char(255);  unique; not null; index: idx_uuid" json:"uuid"`
	Username string    `gorm:"column:user_name" json:"username"`
	IsActive bool      `gorm:"column:is_active; type:boolean" json:"is_active"`
	Role     []Role    `gorm:"many2many:go_user_roles" json:"role"`
}

func (u *User) TableName() string {
	return "go_db_user"
}
