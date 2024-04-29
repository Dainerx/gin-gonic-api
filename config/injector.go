// go:build wireinject
//go:build wireinject
// +build wireinject

package config

import (
	"gin-gonic-api/app/controller"
	"gin-gonic-api/app/repository"
	"gin-gonic-api/app/service"

	"github.com/google/wire"
)

var db = wire.NewSet(ConnectToDB)

// user and role
var userServiceSet = wire.NewSet(service.UserServiceInit,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
)

var userRepoSet = wire.NewSet(repository.UserRepositoryInit,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
)

var userCtrlSet = wire.NewSet(controller.UserControllerInit,
	wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)),
)

var roleRepoSet = wire.NewSet(repository.RoleRepositoryInit,
	wire.Bind(new(repository.RoleRepository), new(*repository.RoleRepositoryImpl)),
)

// project and slide

var projectRepoSet = wire.NewSet(repository.ProjectRepositoryInit,
	wire.Bind(new(repository.ProjectRepository), new(*repository.ProjectRepositoryImpl)),
)

var projectServiceSet = wire.NewSet(service.ProjectServiceInit,
	wire.Bind(new(service.ProjectService), new(*service.ProjectServiceImpl)),
)

var projectCtrlSet = wire.NewSet(controller.ProjectControllerInit,
	wire.Bind(new(controller.ProjectController), new(*controller.ProjectControllerImpl)),
)

var slideRepoSet = wire.NewSet(repository.SlideRepositoryInit,
	wire.Bind(new(repository.SlideRepository), new(*repository.SlideRepositoryImpl)),
)

// model
var modelRepoSet = wire.NewSet(repository.ModelRepositoryInit,
	wire.Bind(new(repository.ModelRepository), new(*repository.ModelRepositoryImpl)),
)

var modelServiceSet = wire.NewSet(service.ModelServiceInit,
	wire.Bind(new(service.ModelService), new(*service.ModelServiceImpl)),
)

var modelCtrlSet = wire.NewSet(controller.ModelControllerInit,
	wire.Bind(new(controller.ModelController), new(*controller.ModelControllerImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, db, userCtrlSet, userServiceSet, userRepoSet, roleRepoSet,
		projectRepoSet, projectServiceSet, projectCtrlSet, slideRepoSet,
		modelRepoSet, modelServiceSet, modelCtrlSet)
	return nil
}
