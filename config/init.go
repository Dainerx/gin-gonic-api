package config

import (
	"gin-gonic-api/app/controller"
	"gin-gonic-api/app/repository"
	"gin-gonic-api/app/service"
)

type Initialization struct {
	userRepo repository.UserRepository
	userSvc  service.UserService
	UserCtrl controller.UserController
	RoleRepo repository.RoleRepository

	projectRepo    repository.ProjectRepository
	projectService service.ProjectService
	ProjectCtrl    controller.ProjectController
	SlideRepo      repository.SlideRepository

	modelRepo    repository.ModelRepository
	modelService service.ModelService
	ModelCtrl    controller.ModelController
}

func NewInitialization(
	// user and role
	userRepo repository.UserRepository,
	userService service.UserService,
	userCtrl controller.UserController,
	roleRepo repository.RoleRepository,
	// project and slide
	projectRepo repository.ProjectRepository,
	projectService service.ProjectService,
	projectCtrl controller.ProjectController,
	slideRepo repository.SlideRepository,
	// model
	modelRepo repository.ModelRepository,
	modelService service.ModelService,
	modelCtrl controller.ModelController,

) *Initialization {
	return &Initialization{
		userRepo: userRepo,
		userSvc:  userService,
		UserCtrl: userCtrl,
		RoleRepo: roleRepo,

		projectRepo:    projectRepo,
		projectService: projectService,
		ProjectCtrl:    projectCtrl,
		SlideRepo:      slideRepo,

		modelRepo:    modelRepo,
		modelService: modelService,
		ModelCtrl:    modelCtrl,
	}
}
