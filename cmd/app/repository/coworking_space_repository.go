package repository

import (
	"learn-rest-api/cmd/app/constant"
	"learn-rest-api/cmd/app/domain/dao"
	"learn-rest-api/cmd/app/exception"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CoworkingSpaceRepository interface {
	FindCoworkingSpaceByID(id int) dao.CoworkingSpace
	SaveCoworkingSpace(m *dao.CoworkingSpace) dao.CoworkingSpace
	FindAllCoworkingSpace() []dao.CoworkingSpace
	DeleteCoworkingSpace(m *dao.CoworkingSpace)
}

type coworkingSpaceRepository struct {
	db *gorm.DB
}

func CoworkingSpaceRepositoryInit(db *gorm.DB) CoworkingSpaceRepository {
	db.AutoMigrate(&dao.CoworkingSpace{})
	return &coworkingSpaceRepository{db: db}
}

// FindAllCoworkingSpace implements CoworkingSpaceRepostory
func (cs *coworkingSpaceRepository) FindAllCoworkingSpace() []dao.CoworkingSpace {
	var coworkingSpaces []dao.CoworkingSpace
	var err = cs.db.Find(&coworkingSpaces).Error
	if err != nil {
		log.Error("Got an error when find all coworking spaces. Error: ", err)
		exception.ThrowNewAppException(constant.UnknownError)
	}

	return coworkingSpaces
}

// FindCoworkingSpaceByID implements CoworkingSpaceRepostory
func (cs *coworkingSpaceRepository) FindCoworkingSpaceByID(id int) dao.CoworkingSpace {
	var coworkingSpace dao.CoworkingSpace
	var err = cs.db.First(&coworkingSpace, id).Error
	if err != nil {
		log.Error("Got an error when find coworking space by id. Error: ", err)
		exception.ThrowNewAppException(constant.DataNotFound)
	}
	return coworkingSpace
}

// SaveCoworkingSpace implements CoworkingSpaceRepostory
func (cs *coworkingSpaceRepository) SaveCoworkingSpace(m *dao.CoworkingSpace) dao.CoworkingSpace {
	var err = cs.db.Save(m).Error
	if err != nil {
		log.Error("Got an error when save coworking space. Error: ", err)
		exception.ThrowNewAppException(constant.UnknownError)
	}
	return *m
}

// DeleteCoworkingSpace implements CoworkingSpaceRepostory
func (cs *coworkingSpaceRepository) DeleteCoworkingSpace(m *dao.CoworkingSpace) {
	var err = cs.db.Delete(m).Error
	if err != nil {
		log.Error("Got an error when delete coworking space. Error: ", err)
		exception.ThrowNewAppException(constant.UnknownError)
	}
}
