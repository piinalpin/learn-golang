package dao

import "learn-rest-api/cmd/app/domain/base"

type CoworkingSpaceRating struct {
	base.BaseModel
	ID               int            `json:"id" gorm:"primary_key; autoIncrement; not null"`
	CoworkingSpaceID int            `json:"-" gorm:"not null"`
	UserID           int            `json:"-" gorm:"not null"`
	RatingScale      int            `json:"rating_scale" gorm:"not null"`
	Comment          string         `json:"comment" gorm:"null; size:500"`
	CoworkingSpace   CoworkingSpace `json:"coworking_space" gorm:"foreignkey:CoworkingSpaceID"`
	User             User           `json:"user" gorm:"foreignkey:UserID"`
}

func (CoworkingSpaceRating) TableName() string {
	return "COWORKING_SPACE_RATING"
}
