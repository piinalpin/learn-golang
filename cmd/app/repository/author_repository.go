package repository

import (
	"learn-rest-api/cmd/app/model"
	"learn-rest-api/config"
	"log"

	"gorm.io/gorm"
)

type AuthorRepository struct {
	Db *gorm.DB
}

func AuthorRepositoryInit() *AuthorRepository {
	var db = config.InitDB()
	db.AutoMigrate(&model.Author{})
	return &AuthorRepository{Db: db}
}

func (repository *AuthorRepository) FindAllAuthors() []model.Author {
	var authors []model.Author
	var err = repository.Db.Find(&authors).Error
	if err != nil {
		log.Panic("Error finding all authors. Error: ", err)
	}

	return authors
}
