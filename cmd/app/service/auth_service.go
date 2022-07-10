package service

import (
	"learn-rest-api/cmd/app/component"
	respkey "learn-rest-api/cmd/app/constant"
	"learn-rest-api/cmd/app/exception"
	"learn-rest-api/cmd/app/form"
	"learn-rest-api/cmd/app/repository"
	"learn-rest-api/cmd/app/validator"
	"learn-rest-api/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type AuthService interface {
	Login(c *gin.Context)
	RefreshToken(c *gin.Context)
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
	defer exception.AppExceptionHandler(c)
	log.Info("Begin execute login")
	var authForm form.LoginForm
	validator.BindJSON(c, &authForm)

	log.Debug("Find user by username: ", authForm.Username)
	user, err := a.userRepo.FindUserByUsername(authForm.Username)

	if err != nil {
		log.Error("Error when find user : ", err)
		exception.ThrowNewAppException(respkey.Unauthorized)
	}

	if !component.Matches(user.Password, authForm.Password) {
		log.Error("Password not match")
		exception.ThrowNewAppException(respkey.Unauthorized)
	}

	token := a.tokenUtil.GenerateAccessToken(user)
	c.JSON(http.StatusOK, pkg.BuildResponse(respkey.Success, token))
}

// RefreshToken implements AuthService
func (a *authService) RefreshToken(c *gin.Context) {
	defer exception.AppExceptionHandler(c)

	var refreshTokenForm form.RefreshTokenForm
	validator.BindJSON(c, &refreshTokenForm)

	token := a.tokenUtil.RefreshToken(refreshTokenForm.RefreshToken)
	c.JSON(http.StatusOK, pkg.BuildResponse(respkey.Success, token))
}