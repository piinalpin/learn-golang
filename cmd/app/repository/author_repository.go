package repository

import (
	"learn-rest-api/cmd/app/domain/dao"

	log "github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

type AuthorRepository interface {
	FindAllAuthors() []dao.Author
	SaveAuthor(m *dao.Author) dao.Author
}

type authorRepository struct {
	db *gorm.DB
}

func AuthorRepositoryInit(db *gorm.DB) AuthorRepository {
	db.AutoMigrate(&dao.Author{})
	return &authorRepository{db: db}
}

// FindAllAuthors implements AuthorRepository
func (a *authorRepository) FindAllAuthors() []dao.Author {
	var authors []dao.Author
	var err = a.db.Find(&authors).Error
	if err != nil {
		log.Error("Got an error finding all authors. Error: ", err)
	}

	return authors
}

// SaveAUthor implements AuthorRepository
func (a *authorRepository) SaveAuthor(m *dao.Author) dao.Author {
	var err = a.db.Save(m).Error
	if err != nil {
		log.Error("Got an error when save author. Error: ", err)
	}
	return *m
}