package config

import (
	"learn-rest-api/cmd/app/controller"
	"learn-rest-api/cmd/app/repository"
	"learn-rest-api/cmd/app/service"

	"gorm.io/gorm"
)

type Initialization struct {
	db         *gorm.DB
	authorRepo repository.AuthorRepository
	authorSvc  service.AuthorService

	AuthorCtrl controller.AuthorController
}

func Init() *Initialization {
	InitLog()
	db := InitDB()
	authorRepo := repository.AuthorRepositoryInit(db)
	authorSvc := service.AuthorServiceInit(authorRepo)
	authorCtrl := controller.AuthorControllerInit(authorSvc)
	return &Initialization{
		db:         db,
		authorRepo: authorRepo,
		authorSvc:  authorSvc,
		AuthorCtrl: authorCtrl,
	}
}
