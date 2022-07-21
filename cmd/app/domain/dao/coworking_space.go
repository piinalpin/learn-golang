package dao

import "learn-rest-api/cmd/app/domain/base"

type CoworkingSpace struct {
	base.BaseModel
	ID         int    `json:"id" gorm:"primary_key; autoIncrement; not null"`
	Name       string `json:"name" gorm:"not null; size:255"`
	Floor      string `json:"floor" gorm:"not null; size:255"`
	Code       string `json:"code" gorm:"null; size:255"`
	Capacity   int    `json:"capacity" gorm:"not null"`
	ProvinceId int    `json:"province_id" gorm:"not null"`
	CityId     int    `json:"city_id" gorm:"not null"`
	DistrictId int    `json:"district_id" gorm:"not null"`
	PostalCode string `json:"postal_code" gorm:"not null; size:255"`
}

func (CoworkingSpace) TableName() string {
	return "COWORKING_SPACE"
}
