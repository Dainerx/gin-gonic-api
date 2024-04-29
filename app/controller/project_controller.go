package controller

import (
	"gin-gonic-api/app/service"

	"github.com/gin-gonic/gin"
)

type ProjectController interface {
	GetAllProjectData(c *gin.Context)
	AddProjectData(c *gin.Context)
	GetProjectById(c *gin.Context)
	UpdateProjectData(c *gin.Context)
	DeleteProject(c *gin.Context)
}

type ProjectControllerImpl struct {
	svc service.ProjectService
}

func (u ProjectControllerImpl) GetAllProjectData(c *gin.Context) {
	u.svc.GetAllProject(c)
}

func (u ProjectControllerImpl) AddProjectData(c *gin.Context) {
	u.svc.AddProjectData(c)
}

func (u ProjectControllerImpl) GetProjectById(c *gin.Context) {
	u.svc.GetProjectById(c)
}

func (u ProjectControllerImpl) UpdateProjectData(c *gin.Context) {
	u.svc.UpdateProjectData(c)
}

func (u ProjectControllerImpl) DeleteProject(c *gin.Context) {
	u.svc.DeleteProject(c)
}

func ProjectControllerInit(projectService service.ProjectService) *ProjectControllerImpl {
	return &ProjectControllerImpl{
		svc: projectService,
	}
}
