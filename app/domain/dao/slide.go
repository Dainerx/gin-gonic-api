package dao

import "gorm.io/gorm"

type Slide struct {
	gorm.Model
	FilePath  string `gorm:"column:file_path" json:"file_path"`
	Thumbnail string `gorm:"column:thumbnail" json:"thumbnail"`
	Metadata  bool   `gorm:"column:metadata" json:"metadata"`
	ProjectID uint
}
