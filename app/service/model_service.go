package service

import (
	"gin-gonic-api/app/constant"
	"gin-gonic-api/app/domain/dao"
	"gin-gonic-api/app/pkg"
	"gin-gonic-api/app/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ModelService interface {
	GetAllModel(c *gin.Context)
	GetModelById(c *gin.Context)
	AddModelData(c *gin.Context)
	UpdateModelData(c *gin.Context)
	DeleteModel(c *gin.Context)
}

type ModelServiceImpl struct {
	modelRepository repository.ModelRepository
}

func (u ModelServiceImpl) UpdateModelData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program update model data by id")
	modelID, _ := strconv.Atoi(c.Param("modelID"))

	var request dao.Model
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := u.modelRepository.FindModelById(modelID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	data.Name = request.Name
	data.Description = request.Description
	data.PrivateScore = request.PrivateScore
	data.CommunnityScore = request.CommunnityScore
	u.modelRepository.Save(&data)

	if err != nil {
		log.Error("Happened error when updating data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u ModelServiceImpl) GetModelById(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get model by id")
	modelID, _ := strconv.Atoi(c.Param("modelID"))

	data, err := u.modelRepository.FindModelById(modelID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u ModelServiceImpl) AddModelData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data model")
	var request dao.Model
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := u.modelRepository.Save(&request)
	if err != nil {
		log.Error("Happened error when saving data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u ModelServiceImpl) GetAllModel(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute get all data model")

	data, err := u.modelRepository.FindAllModel()
	if err != nil {
		log.Error("Happened Error when find all model data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u ModelServiceImpl) DeleteModel(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute delete data model by id")
	modelID, _ := strconv.Atoi(c.Param("modelID"))

	err := u.modelRepository.DeleteModelById(modelID)
	if err != nil {
		log.Error("Happened Error when try delete data model from DB. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func ModelServiceInit(modelRepository repository.ModelRepository) *ModelServiceImpl {
	return &ModelServiceImpl{
		modelRepository: modelRepository,
	}
}
