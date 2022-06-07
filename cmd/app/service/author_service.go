package service

import (
	"learn-rest-api/cmd/app/constant"
	"learn-rest-api/cmd/app/exception"
	"learn-rest-api/cmd/app/form"
	"learn-rest-api/cmd/app/repository"
	"learn-rest-api/cmd/app/validator"
	"learn-rest-api/pkg"
	"log"
	"net/http"

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

	log.Print("Create user")
	var authorForm form.AuthorForm
	validator.BindJSON(c, &authorForm)
	log.Default().Print("Author form: ", authorForm)
	c.JSON(http.StatusOK, pkg.BuildResponse(respkey.Success, pkg.Null()))
}

// GetAllAuthor implements AuthorService
func (a *authorService) GetAllAuthor(c *gin.Context) {
	log.Print("Get all authors")
	var authors = a.repository.FindAllAuthors()
	c.JSON(http.StatusOK, pkg.BuildResponse(respkey.Success, authors))
}