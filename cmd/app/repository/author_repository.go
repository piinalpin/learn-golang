package repository

import (
	"learn-rest-api/cmd/app/model"
	log "github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

type AuthorRepository interface {
	FindAllAuthors() []model.Author
	SaveAuthor(m *model.Author) model.Author
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
		log.Error("Got an error finding all authors. Error: ", err)
	}

	return authors
}

// SaveAUthor implements AuthorRepository
func (a *authorRepository) SaveAuthor(m *model.Author) model.Author {
	var err = a.db.Save(m).Error
	if err != nil {
		log.Error("Got an error when save author. Error: ", err)
	}
	return *m
}