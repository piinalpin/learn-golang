package controller

import (
	"learn-rest-api/cmd/app/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Me(c *gin.Context)
}

type userController struct {
	svc service.UserService
}

func UserControllerInit(s service.UserService) UserController {
	return &userController{
		svc: s,
	}
}

// Me implements UserController
func (u *userController) Me(c *gin.Context) {
	u.svc.Me(c)
}