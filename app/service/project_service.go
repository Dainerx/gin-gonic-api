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

type ProjectService interface {
	GetAllProject(c *gin.Context)
	GetProjectById(c *gin.Context)
	AddProjectData(c *gin.Context)
	UpdateProjectData(c *gin.Context)
	DeleteProject(c *gin.Context)
}

type ProjectServiceImpl struct {
	projectRepository repository.ProjectRepository
}

func (u ProjectServiceImpl) UpdateProjectData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program update project data by id")
	projectID, _ := strconv.Atoi(c.Param("projectID"))

	var request dao.Project
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := u.projectRepository.FindProjectById(projectID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	data.Name = request.Name
	data.Description = request.Description
	// TODO thunmbail and slides
	u.projectRepository.Save(&data)

	if err != nil {
		log.Error("Happened error when updating data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u ProjectServiceImpl) GetProjectById(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get project by id")
	projectID, _ := strconv.Atoi(c.Param("projectID"))

	data, err := u.projectRepository.FindProjectById(projectID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u ProjectServiceImpl) AddProjectData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data project")
	var request dao.Project
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := u.projectRepository.Save(&request)
	if err != nil {
		log.Error("Happened error when saving data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u ProjectServiceImpl) GetAllProject(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute get all data project")

	data, err := u.projectRepository.FindAllProject()
	if err != nil {
		log.Error("Happened Error when find all project data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u ProjectServiceImpl) DeleteProject(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute delete data project by id")
	projectID, _ := strconv.Atoi(c.Param("projectID"))

	err := u.projectRepository.DeleteProjectById(projectID)
	if err != nil {
		log.Error("Happened Error when try delete data project from DB. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func ProjectServiceInit(projectRepository repository.ProjectRepository) *ProjectServiceImpl {
	return &ProjectServiceImpl{
		projectRepository: projectRepository,
	}
}
