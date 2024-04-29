package repository

import (
	"gin-gonic-api/app/domain/dao"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProjectRepository interface {
	FindAllProject() ([]dao.Project, error)
	FindProjectById(id int) (dao.Project, error)
	Save(project *dao.Project) (dao.Project, error)
	DeleteProjectById(id int) error
}

type ProjectRepositoryImpl struct {
	db *gorm.DB
}

func (u ProjectRepositoryImpl) FindAllProject() ([]dao.Project, error) {
	var projects []dao.Project

	var err = u.db.Preload("Slides").Find(&projects).Error
	if err != nil {
		log.Error("Got an error finding all couples. Error: ", err)
		return nil, err
	}

	return projects, nil
}

func (u ProjectRepositoryImpl) FindProjectById(id int) (dao.Project, error) {
	// project := dao.Project{
	// 	ID: id,
	// }
	// err := u.db.Preload("Slides").First(&project).Error
	// if err != nil {
	// 	log.Error("Got and error when find project by id. Error: ", err)
	// 	return dao.Project{}, err
	// }
	return dao.Project{}, nil
}

func (u ProjectRepositoryImpl) Save(project *dao.Project) (dao.Project, error) {
	var err = u.db.Save(project).Error
	if err != nil {
		log.Error("Got an error when save project. Error: ", err)
		return dao.Project{}, err
	}
	return *project, nil
}

func (u ProjectRepositoryImpl) UpdateProjectById() {
	//TODO implement me
	panic("implement me")
}

func (u ProjectRepositoryImpl) DeleteProjectById(id int) error {
	err := u.db.Delete(&dao.Project{}, id).Error
	if err != nil {
		log.Error("Got an error when delete project. Error: ", err)
		return err
	}
	return nil
}

func ProjectRepositoryInit(db *gorm.DB) *ProjectRepositoryImpl {
	// db.AutoMigrate(&dao.Project{})
	return &ProjectRepositoryImpl{
		db: db,
	}
}
