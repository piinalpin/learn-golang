package dao

import "learn-rest-api/cmd/app/domain/base"

type CoworkingSpaceRating struct {
	base.BaseModel
	ID               uint           `json:"id" gorm:"primary_key; autoIncrement; not null"`
	CoworkingSpaceID uint           `json:"-" gorm:"not null"`
	UserID           uint           `json:"-" gorm:"not null"`
	RatingScale      uint           `json:"rating_scale" gorm:"not null"`
	Comment          string         `json:"comment" gorm:"null; size:500"`
	CoworkingSpace   CoworkingSpace `json:"coworking_space" gorm:"foreignkey:CoworkingSpaceID"`
	User             User           `json:"user" gorm:"foreignkey:UserID"`
}

func (CoworkingSpaceRating) TableName() string {
	return "COWORKING_SPACE_RATING"
}
