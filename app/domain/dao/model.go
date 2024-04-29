package dao

import "gorm.io/gorm"

type Model struct {
	gorm.Model
	Name            string  `gorm:"column:name" json:"name"`
	Description     string  `gorm:"column:description" json:"description"`
	Thumbnail       string  `gorm:"column:thumbnail" json:"thumbnail"`
	PrivateScore    float32 `gorm:"column:private_score" json:"private_score"`
	CommunnityScore float32 `gorm:"column:community_score" json:"community_score"`
}
