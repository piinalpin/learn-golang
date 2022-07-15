package repository

import (
	"learn-rest-api/cmd/app/domain/dao"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CoworkingSpaceRepostory interface {
	FindCoworkingSpaceByID(id int) (dao.CoworkingSpace, error)
	SaveCoworkingSpace(m *dao.CoworkingSpace) (dao.CoworkingSpace, error)
	FindAllCoworkingSpace() ([]dao.CoworkingSpace, error)
}

type coworkingSpaceRepository struct {
	db *gorm.DB
}

func CoworkingSpaceRepositoryInit(db *gorm.DB) CoworkingSpaceRepostory {
	db.AutoMigrate(&dao.CoworkingSpace{})
	return &coworkingSpaceRepository{db: db}
}

// FindAllCoworkingSpace implements CoworkingSpaceRepostory
func (cs *coworkingSpaceRepository) FindAllCoworkingSpace() ([]dao.CoworkingSpace, error) {
	var coworkingSpaces []dao.CoworkingSpace
	var err = cs.db.Find(&coworkingSpaces).Error
	if err != nil {
		log.Error("Got an error when find all coworking spaces. Error: ", err)
		return nil, err
	}

	return coworkingSpaces, nil
}

// FindCoworkingSpaceByID implements CoworkingSpaceRepostory
func (cs *coworkingSpaceRepository) FindCoworkingSpaceByID(id int) (dao.CoworkingSpace, error) {
	var coworkingSpace dao.CoworkingSpace
	var err = cs.db.First(&coworkingSpace, id).Error
	if err != nil {
		log.Error("Got an error when find coworking space by id. Error: ", err)
		return dao.CoworkingSpace{}, err
	}
	return coworkingSpace, nil
}

// SaveCoworkingSpace implements CoworkingSpaceRepostory
func (cs *coworkingSpaceRepository) SaveCoworkingSpace(m *dao.CoworkingSpace) (dao.CoworkingSpace, error) {
	var err = cs.db.Save(m).Error
	if err != nil {
		log.Error("Got an error when save coworking space. Error: ", err)
		return dao.CoworkingSpace{}, err
	}
	return *m, nil
}