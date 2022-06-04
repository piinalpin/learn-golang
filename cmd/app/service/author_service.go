package service

import (
	"learn-rest-api/cmd/app/constant"
	"learn-rest-api/cmd/app/repository"
	"learn-rest-api/pkg"
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
)

var authorRepository = repository.AuthorRepositoryInit()

func GetAllAuthor(c *gin.Context) (ctx gin.Context) {
	log.Print("Get all authors")
	var authors = authorRepository.FindAllAuthors()
	var responseKey constant.ResponseKey = constant.Success
	c.JSON(http.StatusOK, pkg.BuildResponse(responseKey, authors))
	return
}