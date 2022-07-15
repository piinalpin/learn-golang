package config

import (
	"learn-rest-api/cmd/app/component"
	"learn-rest-api/cmd/app/controller"
	"learn-rest-api/cmd/app/repository"
	"learn-rest-api/cmd/app/service"

	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

type Initialization struct {
	db         		*gorm.DB
	redis	  		*redis.Client
	
	userRepo   		repository.UserRepository
	userRoleRepo	repository.UserRoleRepository

	TokenUtil 		component.TokenProvider
	sessionStorage	component.SessionStorage

	authSvc   service.AuthService
	userSvc   service.UserService

	AuthCtrl   controller.AuthController
	UserCtrl   controller.UserController
}

func Init() *Initialization {
	InitLog()
	db := InitDB()
	redis := InitRedis()
	userRepo := repository.UserRepositoryInit(db)
	userRoleRepo := repository.UserRoleRepositoryInit(db)

	sessionStorage := component.SessionStorageInit(redis)
	tokenUtil := component.TokenProviderInit(userRepo, sessionStorage)

	authSvc := service.AuthServiceInit(tokenUtil, userRepo)
	userSvc := service.UserServiceInit(userRepo, sessionStorage)

	authCtrl := controller.AuthControllerInit(authSvc)
	userCtrl := controller.UserControllerInit(userSvc)
	return &Initialization{
		db:         	db,
		redis:	  		redis,
		userRepo:   	userRepo,
		userRoleRepo: 	userRoleRepo,

		authSvc:   		authSvc,
		userSvc:   		userSvc,

		TokenUtil: 		tokenUtil,
		sessionStorage: sessionStorage,

		AuthCtrl:   	authCtrl,
		UserCtrl:  		userCtrl,
	}
}
