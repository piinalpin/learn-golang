package model

import (
	"learn-rest-api/cmd/app/model/base"
)

type Author struct {
	base.BaseModel
	ID 				uint 	`json:"id" gorm:"primary_key; autoIncrement; not null"`
	Name			string	`json:"name" gorm:"not null; size:255"`
	IdentityNumber	string	`json:"identity_number" gorm:"not null; size:255"`
}

func (Author) TableName() string {
	return "M_AUTHOR"
}