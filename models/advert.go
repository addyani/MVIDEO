package models

import (
	"gorm.io/gorm"
)

type Advert struct {
	gorm.Model
	Id           int    `form:"id" json:"id" validate:"required"`
	Advert       string `form:"advert" json:"advert" validate:"required"`
	View         int    `form:"view" json:"view" validate:"required"`
	UserIdAdvert int   `gorm:"foreignKey:UserIdUserIdAdvert"`
	IdMIklan     uint
}

func ViewAdvert(db *gorm.DB, advert *[]Advert, id int) (err error) {
	err = db.Where(&Advert{UserIdAdvert: id,}).Find(advert).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadAdsById(db *gorm.DB, advert *Advert, id int) (err error) {
	err = db.Where("id=?", id).First(advert).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateAds(db *gorm.DB, advert *Advert) (err error) {

	

	db.Save(advert)
	
	return nil
}