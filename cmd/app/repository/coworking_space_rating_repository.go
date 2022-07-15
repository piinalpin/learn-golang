package repository

import (
	"learn-rest-api/cmd/app/domain/dao"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CoworkingSpaceRatingRepository interface {
	FindAllCoworkingSpaceRating() ([]dao.CoworkingSpaceRating, error)
	FindAllCoworkingSpaceRatingByCoworkingSpaceID(coworkingSpaceID int) ([]dao.CoworkingSpaceRating, error)
	FindAllCoworkingSpaceRatingByUserID(userID int) ([]dao.CoworkingSpaceRating, error)
	FindCoworkingSpaceRatingByID(id int) (dao.CoworkingSpaceRating, error)
	SaveCoworkingSpaceRating(m *dao.CoworkingSpaceRating) (dao.CoworkingSpaceRating, error)
}

type coworkingSpaceRatingRepository struct {
	db *gorm.DB
}

func CoworkingSpaceRatingRepositoryInit(db *gorm.DB) CoworkingSpaceRatingRepository {
	db.AutoMigrate(&dao.CoworkingSpaceRating{})
	return &coworkingSpaceRatingRepository{db: db}
}

// FindAllCoworkingSpaceRating implements CoworkingSpaceRatingRepository
func (cs *coworkingSpaceRatingRepository) FindAllCoworkingSpaceRating() ([]dao.CoworkingSpaceRating, error) {
	var coworkingSpaceRatings []dao.CoworkingSpaceRating
	var err = cs.db.Find(&coworkingSpaceRatings).Error
	if err != nil {
		log.Error("Got an error when find all coworking space ratings. Error: ", err)
		return nil, err
	}
	
	return coworkingSpaceRatings, nil
}

// FindAllCoworkingSpaceRatingByCoworkingSpaceID implements CoworkingSpaceRatingRepository
func (cs *coworkingSpaceRatingRepository) FindAllCoworkingSpaceRatingByCoworkingSpaceID(coworkingSpaceID int) ([]dao.CoworkingSpaceRating, error) {
	var coworkingSpaceRatings []dao.CoworkingSpaceRating
	var err = cs.db.Model(
		dao.CoworkingSpaceRating {
			CoworkingSpaceID: uint(coworkingSpaceID),
		},
	).Find(&coworkingSpaceRatings).Error
	if err != nil {
		log.Error("Got an error when find all coworking space ratings by coworking space id. Error: ", err)
		return nil, err
	}
	return coworkingSpaceRatings, nil
}

// FindAllCoworkingSpaceRatingByUserID implements CoworkingSpaceRatingRepository
func (cs *coworkingSpaceRatingRepository) FindAllCoworkingSpaceRatingByUserID(userID int) ([]dao.CoworkingSpaceRating, error) {
	var coworkingSpaceRatings []dao.CoworkingSpaceRating
	var err = cs.db.Model(
		dao.CoworkingSpaceRating {
			UserID: uint(userID),
		},
	).Find(&coworkingSpaceRatings).Error
	if err != nil {
		log.Error("Got an error when find all coworking space ratings by user id. Error: ", err)
		return nil, err
	}
	return coworkingSpaceRatings, nil
}

// FindCoworkingSpaceRatingByID implements CoworkingSpaceRatingRepository
func (cs *coworkingSpaceRatingRepository) FindCoworkingSpaceRatingByID(id int) (dao.CoworkingSpaceRating, error) {
	var coworkingSpaceRating dao.CoworkingSpaceRating
	var err = cs.db.First(&coworkingSpaceRating, id).Error
	if err != nil {
		log.Error("Got an error when find coworking space rating by id. Error: ", err)
		return dao.CoworkingSpaceRating{}, err
	}

	return coworkingSpaceRating, nil
}

// SaveCoworkingSpaceRating implements CoworkingSpaceRatingRepository
func (cs *coworkingSpaceRatingRepository) SaveCoworkingSpaceRating(m *dao.CoworkingSpaceRating) (dao.CoworkingSpaceRating, error) {
	var err = cs.db.Save(m).Error
	if err != nil {
		log.Error("Got an error when save coworking space rating. Error: ", err)
		return dao.CoworkingSpaceRating{}, err
	}
	return *m, nil
}