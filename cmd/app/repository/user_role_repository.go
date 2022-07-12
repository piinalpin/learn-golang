package repository

import (
	"learn-rest-api/cmd/app/domain/dao"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRoleRepository interface {
	SaveUserRole(m *dao.UserRole) (dao.UserRole, error)
}

type userRoleRepository struct {
	db *gorm.DB
}

func UserRoleRepositoryInit(db *gorm.DB) UserRoleRepository {
	db.AutoMigrate(&dao.UserRole{})
	return &userRoleRepository{db: db}
}

// SaveUserRole implements UserRoleRepository
func (u *userRoleRepository) SaveUserRole(m *dao.UserRole) (dao.UserRole, error) {
	var err = u.db.Save(m).Error
	if err != nil {
		log.Error("Got an error when save user role. Error: ", err)
		return dao.UserRole{}, err
	}
	return *m, nil
}