package service

import (
	"learn-rest-api/cmd/app/repository"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type UserService interface {
	Me(c *gin.Context)
}

type userService struct {
	repository repository.UserRepository
}

func UserServiceInit(r repository.UserRepository) UserService {
	return &userService{
		repository: r,
	}
}

// Me implements UserService
func (*userService) Me(c *gin.Context) {
	log.Info("unimplemented")
}
