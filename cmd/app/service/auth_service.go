package service

import (
	"learn-rest-api/cmd/app/component"
	respkey "learn-rest-api/cmd/app/constant"
	"learn-rest-api/cmd/app/repository"
	"learn-rest-api/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type AuthService interface {
	Login(c *gin.Context)
}

type authService struct {
	tokenUtil component.TokenProvider
	userRepo  repository.UserRepository
}

func AuthServiceInit(t component.TokenProvider, u repository.UserRepository) AuthService {
	return &authService{
		tokenUtil: t,
		userRepo:  u,
	}
}

// Login implements AuthService
func (a *authService) Login(c *gin.Context) {
	log.Info("Begin login")

	token := a.tokenUtil.GenerateAccessToken("anyuser")
	c.JSON(http.StatusOK, pkg.BuildResponse(respkey.Success, token))
}