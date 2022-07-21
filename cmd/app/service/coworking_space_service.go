package service

import (
	"learn-rest-api/cmd/app/constant"
	"learn-rest-api/cmd/app/domain/dao"
	"learn-rest-api/cmd/app/domain/dto"
	"learn-rest-api/cmd/app/exception"
	"learn-rest-api/cmd/app/repository"
	"learn-rest-api/cmd/app/validator"
	"learn-rest-api/pkg"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type CoworkingSpaceService interface {
	GetAllCoworkingSpace(c *gin.Context)
	GetCoworkingSpaceById(c *gin.Context)
	SaveCoworkingSpace(c *gin.Context)
	UpdateCoworkingSpace(c *gin.Context)
	DeleteCoworkingSpace(c *gin.Context)
}

type coworkingSpaceService struct {
	coworkingSpaceRepository repository.CoworkingSpaceRepository
}

func CoworkingSpaceServiceInit(coworkingSpaceRepository repository.CoworkingSpaceRepository) CoworkingSpaceService {
	return &coworkingSpaceService{coworkingSpaceRepository: coworkingSpaceRepository}
}

// FindAllCoworkingSpace implements CoworkingSpaceService
func (cs *coworkingSpaceService) GetAllCoworkingSpace(c *gin.Context) {
	defer exception.AppExceptionHandler(c)

	log.Info("Get all coworking spaces")
	coworkingSpaces := cs.coworkingSpaceRepository.FindAllCoworkingSpace()

	response, _ := pkg.TypeConverter(&coworkingSpaces, &[]dto.CoworkingSpaceDto{})

	log.Info("Coworking spaces size: ", len(*response))
	c.JSON(200, pkg.BuildResponse(constant.Success, response))
}

// FindCoworkingSpaceById implements CoworkingSpaceService
func (cs *coworkingSpaceService) GetCoworkingSpaceById(c *gin.Context) {
	defer exception.AppExceptionHandler(c)

	log.Info("Get coworking space by id")
	id, _ := strconv.Atoi(c.Param("id"))

	coworkingSpace := cs.coworkingSpaceRepository.FindCoworkingSpaceByID(id)
	if &coworkingSpace == nil {
		log.Error("Coworking space not found")
		exception.ThrowNewAppException_(constant.DataNotFound.GetKey(), "Coworking space not found")
	}

	response, _ := pkg.TypeConverter(&coworkingSpace, &dto.CoworkingSpaceDto{})
	c.JSON(200, pkg.BuildResponse(constant.Success, response))
}

// SaveCoworkingSpace implements CoworkingSpaceService
func (cs *coworkingSpaceService) SaveCoworkingSpace(c *gin.Context) {
	defer exception.AppExceptionHandler(c)

	log.Info("Save coworking space")
	var coworkingSpaceDto dto.CoworkingSpaceDto
	validator.BindJSON(c, &coworkingSpaceDto)

	coworkingSpace, _ := pkg.TypeConverter(&coworkingSpaceDto, &dao.CoworkingSpace{})

	log.Info("Saving coworking space")
	cs.coworkingSpaceRepository.SaveCoworkingSpace(coworkingSpace)

	response, _ := pkg.TypeConverter(&coworkingSpace, &dto.CoworkingSpaceDto{})
	c.JSON(200, pkg.BuildResponse(constant.Success, response))
}

// UpdateCoworkingSpace implements CoworkingSpaceService
func (cs *coworkingSpaceService) UpdateCoworkingSpace(c *gin.Context) {
	defer exception.AppExceptionHandler(c)

	log.Info("Update coworking space")
	id, _ := strconv.Atoi(c.Param("id"))

	var coworkingSpaceDto dto.CoworkingSpaceDto
	validator.BindJSON(c, &coworkingSpaceDto)

	coworkingSpace := cs.coworkingSpaceRepository.FindCoworkingSpaceByID(id)
	if &coworkingSpace == nil {
		log.Error("Coworking space not found")
		exception.ThrowNewAppException_(constant.DataNotFound.GetKey(), "Coworking space not found")
	}

	coworkingSpace.Name = coworkingSpaceDto.Name
	coworkingSpace.Floor = coworkingSpaceDto.Floor
	coworkingSpace.Capacity = coworkingSpaceDto.Capacity
	coworkingSpace.Code = coworkingSpaceDto.Code
	coworkingSpace.ProvinceId = coworkingSpaceDto.ProvinceId
	coworkingSpace.CityId = coworkingSpaceDto.CityId
	coworkingSpace.DistrictId = coworkingSpaceDto.DistrictId
	coworkingSpace.PostalCode = coworkingSpaceDto.PostalCode
	cs.coworkingSpaceRepository.SaveCoworkingSpace(&coworkingSpace)

	response, _ := pkg.TypeConverter(&coworkingSpace, &dto.CoworkingSpaceDto{})
	c.JSON(200, pkg.BuildResponse(constant.Success, response))
}

// DeleteCoworkingSpace implements CoworkingSpaceService
func (cs *coworkingSpaceService) DeleteCoworkingSpace(c *gin.Context) {
	defer exception.AppExceptionHandler(c)

	log.Info("Delete coworking space")
	id, _ := strconv.Atoi(c.Param("id"))

	coworkingSpace := cs.coworkingSpaceRepository.FindCoworkingSpaceByID(id)
	if &coworkingSpace == nil {
		log.Error("Coworking space not found")
		exception.ThrowNewAppException_(constant.DataNotFound.GetKey(), "Coworking space not found")
	}

	cs.coworkingSpaceRepository.DeleteCoworkingSpace(&coworkingSpace)
	c.JSON(200, pkg.BuildResponse(constant.Success, pkg.Null()))
}
