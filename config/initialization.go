package config

import (
	"learn-rest-api/cmd/app/component"
	"learn-rest-api/cmd/app/controller"
	"learn-rest-api/cmd/app/repository"
	"learn-rest-api/cmd/app/service"

	"gorm.io/gorm"
)

type Initialization struct {
	db         *gorm.DB
	authorRepo repository.AuthorRepository
	userRepo   repository.UserRepository

	TokenUtil component.TokenProvider

	authSvc   service.AuthService
	authorSvc service.AuthorService
	userSvc   service.UserService

	AuthCtrl   controller.AuthController
	AuthorCtrl controller.AuthorController
}

func Init() *Initialization {
	InitLog()
	db := InitDB()
	authorRepo := repository.AuthorRepositoryInit(db)
	userRepo := repository.UserRepositoryInit(db)

	tokenUtil := component.TokenProviderInit(userRepo)

	authSvc := service.AuthServiceInit(tokenUtil, userRepo)
	authorSvc := service.AuthorServiceInit(authorRepo)
	userSvc := service.UserServiceInit(userRepo)

	authCtrl := controller.AuthControllerInit(authSvc)
	authorCtrl := controller.AuthorControllerInit(authorSvc)
	return &Initialization{
		db:         db,
		authorRepo: authorRepo,
		userRepo:   userRepo,

		authSvc:   authSvc,
		authorSvc: authorSvc,
		userSvc:   userSvc,

		TokenUtil: tokenUtil,

		AuthCtrl:   authCtrl,
		AuthorCtrl: authorCtrl,
	}
}
