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
	
	authorRepo 		repository.AuthorRepository
	userRepo   		repository.UserRepository
	userRoleRepo	repository.UserRoleRepository

	TokenUtil 		component.TokenProvider
	sessionStorage	component.SessionStorage

	authSvc   service.AuthService
	authorSvc service.AuthorService
	userSvc   service.UserService

	AuthCtrl   controller.AuthController
	AuthorCtrl controller.AuthorController
	UserCtrl   controller.UserController
}

func Init() *Initialization {
	InitLog()
	db := InitDB()
	redis := InitRedis()
	authorRepo := repository.AuthorRepositoryInit(db)
	userRepo := repository.UserRepositoryInit(db)
	userRoleRepo := repository.UserRoleRepositoryInit(db)

	sessionStorage := component.SessionStorageInit(redis)
	tokenUtil := component.TokenProviderInit(userRepo, sessionStorage)

	authSvc := service.AuthServiceInit(tokenUtil, userRepo)
	authorSvc := service.AuthorServiceInit(authorRepo)
	userSvc := service.UserServiceInit(userRepo, sessionStorage)

	authCtrl := controller.AuthControllerInit(authSvc)
	authorCtrl := controller.AuthorControllerInit(authorSvc)
	userCtrl := controller.UserControllerInit(userSvc)
	return &Initialization{
		db:         	db,
		redis:	  		redis,
		authorRepo: 	authorRepo,
		userRepo:   	userRepo,
		userRoleRepo: 	userRoleRepo,

		authSvc:   		authSvc,
		authorSvc: 		authorSvc,
		userSvc:   		userSvc,

		TokenUtil: 		tokenUtil,
		sessionStorage: sessionStorage,

		AuthCtrl:   	authCtrl,
		AuthorCtrl: 	authorCtrl,
		UserCtrl:  		userCtrl,
	}
}
