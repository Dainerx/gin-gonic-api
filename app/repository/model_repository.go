package repository

import (
	"gin-gonic-api/app/domain/dao"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ModelRepository interface {
	FindAllModel() ([]dao.Model, error)
	FindModelById(id int) (dao.Model, error)
	Save(model *dao.Model) (dao.Model, error)
	DeleteModelById(id int) error
}

type ModelRepositoryImpl struct {
	db *gorm.DB
}

func (u ModelRepositoryImpl) FindAllModel() ([]dao.Model, error) {
	var models []dao.Model

	var err = u.db.Find(&models).Error
	if err != nil {
		log.Error("Got an error finding all couples. Error: ", err)
		return nil, err
	}

	return models, nil
}

func (u ModelRepositoryImpl) FindModelById(id int) (dao.Model, error) {
	// model := dao.Model{
	// 	ID: id,
	// }
	// err := u.db.Preload("Slides").First(&model).Error
	// if err != nil {
	// 	log.Error("Got and error when find model by id. Error: ", err)
	// 	return dao.Model{}, err
	// }
	return dao.Model{}, nil
}

func (u ModelRepositoryImpl) Save(model *dao.Model) (dao.Model, error) {
	var err = u.db.Save(model).Error
	if err != nil {
		log.Error("Got an error when save model. Error: ", err)
		return dao.Model{}, err
	}
	return *model, nil
}

func (u ModelRepositoryImpl) UpdateModelById() {
	//TODO implement me
	panic("implement me")
}

func (u ModelRepositoryImpl) DeleteModelById(id int) error {
	err := u.db.Delete(&dao.Model{}, id).Error
	if err != nil {
		log.Error("Got an error when delete model. Error: ", err)
		return err
	}
	return nil
}

func ModelRepositoryInit(db *gorm.DB) *ModelRepositoryImpl {
	db.AutoMigrate(&dao.Model{})
	return &ModelRepositoryImpl{
		db: db,
	}
}
