package controller

import (
	"gin-gonic-api/app/service"

	"github.com/gin-gonic/gin"
)

type ModelController interface {
	GetAllModelData(c *gin.Context)
	AddModelData(c *gin.Context)
	GetModelById(c *gin.Context)
	UpdateModelData(c *gin.Context)
	DeleteModel(c *gin.Context)
}

type ModelControllerImpl struct {
	svc service.ModelService
}

func (u ModelControllerImpl) GetAllModelData(c *gin.Context) {
	u.svc.GetAllModel(c)
}

func (u ModelControllerImpl) AddModelData(c *gin.Context) {
	u.svc.AddModelData(c)
}

func (u ModelControllerImpl) GetModelById(c *gin.Context) {
	u.svc.GetModelById(c)
}

func (u ModelControllerImpl) UpdateModelData(c *gin.Context) {
	u.svc.UpdateModelData(c)
}

func (u ModelControllerImpl) DeleteModel(c *gin.Context) {
	u.svc.DeleteModel(c)
}

func ModelControllerInit(modelService service.ModelService) *ModelControllerImpl {
	return &ModelControllerImpl{
		svc: modelService,
	}
}
