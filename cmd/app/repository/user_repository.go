package repository

import (
	"learn-rest-api/cmd/app/domain/dao"

	log "github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserByUsername(username string) (dao.User, error)
	SaveUser(m *dao.User) (dao.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func UserRepositoryInit(db *gorm.DB) UserRepository {
	db.AutoMigrate(&dao.User{})
	return &userRepository{db: db}
}

// FindUserByUsername implements UserRepository
func (u *userRepository) FindUserByUsername(username string) (dao.User, error) {
	var user dao.User
	var err = u.db.Preload("Roles").Where("username = ?", username).First(&user).Error
	if err != nil {
		log.Error("Got an error finding user by username. Error: ", err)
		return dao.User{}, err
	}

	return user, nil
}

// SaveUser implements UserRepository
func (u *userRepository) SaveUser(m *dao.User) (dao.User, error) {
	var err = u.db.Save(m).Error
	if err != nil {
		log.Error("Got an error when save user. Error: ", err)
		return dao.User{}, err
	}
	return *m, nil
}