package models

import (
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Id          int    `form:"id" json:"id" validate:"required"`
	Video       string `form:"video" json:"video" validate:"required"`
	Title       string `form:"title" json:"title" validate:"required"`
	Description string `form:"description" json:"description" validate:"required"`
	DisableAds  bool   `gorm:"default:false"`
	UserIdVideo uint   `gorm:"foreignKey:UserIdUserIdVideo"`
}
