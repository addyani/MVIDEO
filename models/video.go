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

func CreateVideo(db *gorm.DB, newVideo *Video, userId uint) (err error) {
	newVideo.UserIdVideo = userId
	err = db.Create(newCart).Error
	if err != nil {
		return err
	}
	return nil

func ViewVideo(db *gorm.DB, video *Video, id int) (err error) {
	err = db.Where("user_id=?", id).ind(video).Error
	if err != nil {
		return err
	}
	return nil
	}

func FindVideoById(db *gorm.DB, video *Video, id int) (err error) {
	err = db.Where("user_id=?", id).First(video).Error
	if err != nil {
		return err
	}
	return nil
	}
func UpdateVideo(db *gorm.DB, video *Video) (err error) {
	db.Save(video)

	return nil
	}
func DeleteVideo(db *gorm.DB, video *Video, id int) (err error) {
	db.Where("id=?", id).Delete(video)

	return nil
	}