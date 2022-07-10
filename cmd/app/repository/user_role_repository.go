package repository

import (
	"learn-rest-api/cmd/app/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRoleRepository interface {
	SaveUserRole(m *model.UserRole) (model.UserRole, error)
}

type userRoleRepository struct {
	db *gorm.DB
}

func UserRoleRepositoryInit(db *gorm.DB) UserRoleRepository {
	db.AutoMigrate(&model.UserRole{})
	return &userRoleRepository{db: db}
}

// SaveUserRole implements UserRoleRepository
func (u *userRoleRepository) SaveUserRole(m *model.UserRole) (model.UserRole, error) {
	var err = u.db.Save(m).Error
	if err != nil {
		log.Error("Got an error when save user role. Error: ", err)
		return model.UserRole{}, err
	}
	return *m, nil
}