package dao

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string  `gorm:"column:name" json:"name"`
	Description string  `gorm:"column:description" json:"description"`
	Thumbnail   string  `gorm:"column:thumbnail" json:"thumbnail"`
	Slides      []Slide `gorm:"foreignkey:ProjectID" json:"slides"`
}
