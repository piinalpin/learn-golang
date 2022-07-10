package controller

import (
	"learn-rest-api/cmd/app/service"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(c *gin.Context)
	RefreshToken(c *gin.Context)
}

type authController struct {
	svc service.AuthService
}

func AuthControllerInit(s service.AuthService) AuthController {
	return &authController{
		svc: s,
	}
}

// Login implements AuthController
func (a *authController) Login(c *gin.Context) {
	a.svc.Login(c)
}

// RefreshToken implements AuthController
func (a *authController) RefreshToken(c *gin.Context) {
	a.svc.RefreshToken(c)
}