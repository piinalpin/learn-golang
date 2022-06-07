package repository

import (
	"learn-rest-api/cmd/app/model"
	"log"

	"gorm.io/gorm"
)

type AuthorRepository interface {
	FindAllAuthors() []model.Author
}

type authorRepository struct {
	db *gorm.DB
}

func AuthorRepositoryInit(db *gorm.DB) AuthorRepository {
	db.AutoMigrate(&model.Author{})
	return &authorRepository{db: db}
}

// FindAllAuthors implements AuthorRepository
func (a *authorRepository) FindAllAuthors() []model.Author {
	var authors []model.Author
	var err = a.db.Find(&authors).Error
	if err != nil {
		log.Panic("Error finding all authors. Error: ", err)
	}

	return authors
}
