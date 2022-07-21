package dao

import "learn-rest-api/cmd/app/domain/base"

type User struct {
	base.BaseModel
	ID 				int 		`json:"id" gorm:"primary_key; autoIncrement; not null"`
	Username		string		`json:"username" gorm:"not null; size:255; index:,unique"`
	Password		string		`json:"password" gorm:"not null; size:255"`
	Status			string		`json:"status" gorm:"not null; size:255"`
	Roles			[]UserRole	`json:"user_role" gorm:"foreignkey:UserID"`
}

func (User) TableName() string {
	return "M_USER"
}