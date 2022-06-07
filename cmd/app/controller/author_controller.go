package controller

import (
	_ "learn-rest-api/cmd/app/constant"
	"learn-rest-api/cmd/app/service"
	_ "learn-rest-api/pkg"
	_ "net/http"

	"github.com/gin-gonic/gin"
)

type AuthorController interface {
	GetAllAuthor(c *gin.Context)
	CreateAuthor(c *gin.Context)
}

type authorController struct {
	svc service.AuthorService
}

func AuthorControllerInit(s service.AuthorService) AuthorController {
	return &authorController{
		svc: s,
	}
}

// CreateAuthor implements AuthorController
func (a *authorController) CreateAuthor(c *gin.Context) {
	a.svc.CreateAuthor(c)
}

// GetAllAuthor implements AuthorController
func (a *authorController) GetAllAuthor(c *gin.Context) {
	a.svc.GetAllAuthor(c)
}