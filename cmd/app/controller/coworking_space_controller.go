package controller

import (
	"learn-rest-api/cmd/app/service"

	"github.com/gin-gonic/gin"
)

type CoworkingSpaceController interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type coworkingSpaceController struct {
	svc service.CoworkingSpaceService
}

func CoworkingSpaceControllerInit(s service.CoworkingSpaceService) CoworkingSpaceController {
	return &coworkingSpaceController{
		svc: s,
	}
}

// Create implements CoworkingSpaceController
func (csc *coworkingSpaceController) Create(c *gin.Context) {
	csc.svc.SaveCoworkingSpace(c)
}

// Delete implements CoworkingSpaceController
func (csc *coworkingSpaceController) Delete(c *gin.Context) {
	csc.svc.DeleteCoworkingSpace(c)
}

// GetAll implements CoworkingSpaceController
func (csc *coworkingSpaceController) GetAll(c *gin.Context) {
	csc.svc.GetAllCoworkingSpace(c)
}

// GetById implements CoworkingSpaceController
func (csc *coworkingSpaceController) GetById(c *gin.Context) {
	csc.svc.GetCoworkingSpaceById(c)
}

// Update implements CoworkingSpaceController
func (csc *coworkingSpaceController) Update(c *gin.Context) {
	csc.svc.UpdateCoworkingSpace(c)
}
