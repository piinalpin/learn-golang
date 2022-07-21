package repository

import (
	"learn-rest-api/cmd/app/constant"
	"learn-rest-api/cmd/app/domain/dao"
	"learn-rest-api/cmd/app/exception"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CoworkingSpaceRatingRepository interface {
	FindAllCoworkingSpaceRating() []dao.CoworkingSpaceRating
	FindAllCoworkingSpaceRatingByCoworkingSpaceID(coworkingSpaceID int) []dao.CoworkingSpaceRating
	FindAllCoworkingSpaceRatingByUserID(userID int) []dao.CoworkingSpaceRating
	FindCoworkingSpaceRatingByID(id int) dao.CoworkingSpaceRating
	SaveCoworkingSpaceRating(m *dao.CoworkingSpaceRating) dao.CoworkingSpaceRating
}

type coworkingSpaceRatingRepository struct {
	db *gorm.DB
}

func CoworkingSpaceRatingRepositoryInit(db *gorm.DB) CoworkingSpaceRatingRepository {
	db.AutoMigrate(&dao.CoworkingSpaceRating{})
	return &coworkingSpaceRatingRepository{db: db}
}

// FindAllCoworkingSpaceRating implements CoworkingSpaceRatingRepository
func (cs *coworkingSpaceRatingRepository) FindAllCoworkingSpaceRating() []dao.CoworkingSpaceRating {
	var coworkingSpaceRatings []dao.CoworkingSpaceRating
	var err = cs.db.Find(&coworkingSpaceRatings).Error
	if err != nil {
		log.Error("Got an error when find all coworking space ratings. Error: ", err)
		exception.ThrowNewAppException(constant.UnknownError)
	}
	
	return coworkingSpaceRatings
}

// FindAllCoworkingSpaceRatingByCoworkingSpaceID implements CoworkingSpaceRatingRepository
func (cs *coworkingSpaceRatingRepository) FindAllCoworkingSpaceRatingByCoworkingSpaceID(coworkingSpaceID int) []dao.CoworkingSpaceRating {
	var coworkingSpaceRatings []dao.CoworkingSpaceRating
	var err = cs.db.Model(
		dao.CoworkingSpaceRating {
			CoworkingSpaceID: coworkingSpaceID,
		},
	).Find(&coworkingSpaceRatings).Error
	if err != nil {
		log.Error("Got an error when find all coworking space ratings by coworking space id. Error: ", err)
		exception.ThrowNewAppException(constant.UnknownError)
	}
	return coworkingSpaceRatings
}

// FindAllCoworkingSpaceRatingByUserID implements CoworkingSpaceRatingRepository
func (cs *coworkingSpaceRatingRepository) FindAllCoworkingSpaceRatingByUserID(userID int) []dao.CoworkingSpaceRating {
	var coworkingSpaceRatings []dao.CoworkingSpaceRating
	var err = cs.db.Model(
		dao.CoworkingSpaceRating {
			UserID: userID,
		},
	).Find(&coworkingSpaceRatings).Error
	if err != nil {
		log.Error("Got an error when find all coworking space ratings by user id. Error: ", err)
		exception.ThrowNewAppException(constant.UnknownError)
	}
	return coworkingSpaceRatings
}

// FindCoworkingSpaceRatingByID implements CoworkingSpaceRatingRepository
func (cs *coworkingSpaceRatingRepository) FindCoworkingSpaceRatingByID(id int) dao.CoworkingSpaceRating {
	var coworkingSpaceRating dao.CoworkingSpaceRating
	var err = cs.db.First(&coworkingSpaceRating, id).Error
	if err != nil {
		log.Error("Got an error when find coworking space rating by id. Error: ", err)
		exception.ThrowNewAppException(constant.UnknownError)
	}

	return coworkingSpaceRating
}

// SaveCoworkingSpaceRating implements CoworkingSpaceRatingRepository
func (cs *coworkingSpaceRatingRepository) SaveCoworkingSpaceRating(m *dao.CoworkingSpaceRating) dao.CoworkingSpaceRating {
	var err = cs.db.Save(m).Error
	if err != nil {
		log.Error("Got an error when save coworking space rating. Error: ", err)
		exception.ThrowNewAppException(constant.UnknownError)
	}
	return *m
}