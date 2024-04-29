package repository

import (
	"gin-gonic-api/app/domain/dao"

	"gorm.io/gorm"
)

type SlideRepository interface {
	FindAllSlides()
}

type SlideRepositoryImpl struct {
	db *gorm.DB
}

func (r SlideRepositoryImpl) FindAllSlides() {
	panic("implement me")
}

func SlideRepositoryInit(db *gorm.DB) *SlideRepositoryImpl {
	db.AutoMigrate(&dao.Slide{}, &dao.Project{})
	return &SlideRepositoryImpl{
		db: db,
	}
}
