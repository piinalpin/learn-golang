package service

import (
	"learn-rest-api/cmd/app/component"
	"learn-rest-api/cmd/app/constant"
	"learn-rest-api/cmd/app/domain/dao"
	"learn-rest-api/cmd/app/domain/dto"
	"learn-rest-api/cmd/app/exception"
	"learn-rest-api/cmd/app/repository"
	"learn-rest-api/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type UserService interface {
	Me(c *gin.Context)
}

type userService struct {
	repository     repository.UserRepository
	sessionStorage component.SessionStorage
}

func UserServiceInit(r repository.UserRepository, s component.SessionStorage) UserService {
	return &userService{
		repository:     r,
		sessionStorage: s,
	}
}

// Me implements UserService
func (u *userService) Me(c *gin.Context) {
	log.Info("Get user me")
	var tokenMetadata dao.TokenMetadata
	claims, isExists := c.Get(constant.JwtClaims.GetContextKey())

	if !isExists {
		log.Info("Failed get token claims from context")
		exception.ThrowNewAppException(constant.Unauthorized)
	}

	log.Debug("Token claims from context: ", claims)
	err := u.sessionStorage.GetCache(constant.UserSession.GetCacheName(), claims.(*component.JwtCustomClaims).Uuid, &tokenMetadata)

	if err != nil {
		log.Info("User session not found!")
		exception.ThrowNewAppException(constant.Unauthorized)
	}

	userDto, _ := pkg.TypeConverter(&tokenMetadata.User, &dto.UserDto{})
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, userDto))
}
