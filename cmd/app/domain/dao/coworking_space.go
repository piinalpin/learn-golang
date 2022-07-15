package dao

import "learn-rest-api/cmd/app/domain/base"

type CoworkingSpace struct {
	base.BaseModel
	ID         uint   `json:"id" gorm:"primary_key; autoIncrement; not null"`
	Name       string `json:"name" gorm:"not null; size:255"`
	Floor      string `json:"floor" gorm:"not null; size:255"`
	Code       string `json:"code" gorm:"null; size:255"`
	Capacity   uint   `json:"capacity" gorm:"not null"`
	ProvinceId uint   `json:"province_id" gorm:"not null"`
	CityId     uint   `json:"city_id" gorm:"not null"`
	DistrictId uint   `json:"district_id" gorm:"not null"`
	PostalCode string `json:"postal_code" gorm:"not null; size:255"`
}

func (CoworkingSpace) TableName() string {
	return "COWORKING_SPACE"
}
