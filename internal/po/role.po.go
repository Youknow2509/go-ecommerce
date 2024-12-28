package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int64  `gorm:"column:id; type:int;  unique; not null; primaryKey; autoIncrement; commment: 'Primary key is id'" json:"id"`
	RoleName string `gorm:"column:role_name' json:"role_name"`
	RoleNote string `gorm:"column:role_note; type:text;" json:"role_note"`
}

func (r *Role) TableName() string {
	return "go_db_role"
}
