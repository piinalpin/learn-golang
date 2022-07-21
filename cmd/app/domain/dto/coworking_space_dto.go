package dto

import "learn-rest-api/cmd/app/domain/base"

type CoworkingSpaceDto struct {
	base.BaseDto
	ID         int   `json:"id"`
	Name       string `json:"name"`
	Floor      string `json:"floor"`
	Code       string `json:"code"`
	Capacity   int   `json:"capacity"`
	ProvinceId int   `json:"province_id"`
	CityId     int   `json:"city_id"`
	DistrictId int   `json:"district_id"`
	PostalCode string `json:"postal_code"`
}
