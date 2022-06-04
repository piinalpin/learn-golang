package controller

import (
	_ "learn-rest-api/cmd/app/constant"
	"learn-rest-api/cmd/app/service"
	_ "learn-rest-api/pkg"
	_ "net/http"

	"github.com/gin-gonic/gin"
)

type AuthorController struct {}

func (a AuthorController) GetAllAuthor(c *gin.Context) {
	service.GetAllAuthor(c)
}