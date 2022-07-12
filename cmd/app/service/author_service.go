package service

import (
	respkey "learn-rest-api/cmd/app/constant"
	"learn-rest-api/cmd/app/domain/dao"
	"learn-rest-api/cmd/app/exception"
	"learn-rest-api/cmd/app/form"
	"learn-rest-api/cmd/app/repository"
	"learn-rest-api/cmd/app/validator"
	"learn-rest-api/pkg"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type AuthorService interface {
	GetAllAuthor(c *gin.Context)
	CreateAuthor(c *gin.Context)
}

type authorService struct {
	repository repository.AuthorRepository
}

func AuthorServiceInit(r repository.AuthorRepository) AuthorService {
	return &authorService{
		repository: r,
	}
}

// CreateAuthor implements AuthorService
func (a *authorService) CreateAuthor(c *gin.Context) {
	defer exception.AppExceptionHandler(c)

	log.Info("Create author")
	var authorForm form.AuthorForm
	validator.BindJSON(c, &authorForm)
	
	log.Debug("Author form:: ", authorForm)
	author, _ := pkg.TypeConverter[dao.Author](&authorForm)

	log.Info("Saving author")
	a.repository.SaveAuthor(author)
	c.JSON(http.StatusOK, pkg.BuildResponse(respkey.Success, author))
}

// GetAllAuthor implements AuthorService
func (a *authorService) GetAllAuthor(c *gin.Context) {
	log.Info("Get all author")
	var authors = a.repository.FindAllAuthors()
	log.Info("Authors size: ", len(authors))
	c.JSON(http.StatusOK, pkg.BuildResponse(respkey.Success, authors))
}