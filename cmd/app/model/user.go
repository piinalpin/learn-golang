package model

import (
	"encoding/json"
	"learn-rest-api/cmd/app/model/base"
)

type User struct {
	base.BaseModel
	ID 				uint 		`json:"id" gorm:"primary_key; autoIncrement; not null"`
	Username		string		`json:"username" gorm:"not null; size:255"`
	Password		string		`json:"password" gorm:"not null; size:255"`
	Status			string		`json:"status" gorm:"not null; size:255"`
	Roles			[]UserRole	`json:"user_role" gorm:"foreignkey:UserID"`
}

func (User) TableName() string {
	return "M_USER"
}

func (u User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}