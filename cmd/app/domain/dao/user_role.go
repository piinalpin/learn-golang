package dao

import (
	"learn-rest-api/cmd/app/domain/base"
)

type UserRole struct {
	base.BaseModel
	ID 				int 	`json:"id" gorm:"primary_key; autoIncrement; not null"`
	UserID			int	`json:"user_id" gorm:"not null"`
	Role			string	`json:"role" gorm:"not null"`
	User			User	`json:"-" gorm:"foreignkey:UserID"`
}

func (UserRole) TableName() string {
	return "M_USER_ROLE"
}