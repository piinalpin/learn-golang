package dao

import (
	"learn-rest-api/cmd/app/domain/base"
)

type Author struct {
	base.BaseDao
	ID uint `json:"id" gorm:"primary_key; autoIncrement; not null"`
	Name string `json:"name" gorm:"not null; size:255"`
}

func (Author) TableName() string {
	return "M_AUTHOR"
}