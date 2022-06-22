package repository

import (
	"learn-rest-api/cmd/app/model"

	log "github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserByUsername(username string) (model.User, error)
	SaveUser(m *model.User) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func UserRepositoryInit(db *gorm.DB) UserRepository {
	db.AutoMigrate(&model.User{})
	return &userRepository{db: db}
}

// FindUserByUsername implements UserRepository
func (u *userRepository) FindUserByUsername(username string) (model.User, error) {
	var user model.User
	var err = u.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		log.Error("Got an error finding user by username. Error: ", err)
		return model.User{}, err
	}

	return user, nil
}

// SaveUser implements UserRepository
func (u *userRepository) SaveUser(m *model.User) (model.User, error) {
	var err = u.db.Save(m).Error
	if err != nil {
		log.Error("Got an error when save user. Error: ", err)
		return model.User{}, err
	}
	return *m, nil
}