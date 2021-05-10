package models

import (
	"github.com/jinzhu/gorm"
)

type Gateway struct {
	gorm.Model
	Name      string `gorm:"not null" json:"name"`
	AccessKey string `gorm:"not null" json:"accesskey"`
	SecretKey string `gorm:"not null" json:"secretkey"`
	Token     string `gorm:"not null" json:"token"`
	Others    string `gorm:"not null" json:"others"`
	Url       string `gorm:"not null" json:"url"`
	Active    bool   `gorm:"not null" json:"active"`
}


func (data *Gateway) Save(db *gorm.DB) (*Gateway, error) {
	var err error
	err = db.Model(&Gateway{}).Create(&data).Error
	if err != nil {
		return &Gateway{}, err
	}
	return data, nil
}

func (data *Gateway) FindAll(db *gorm.DB) (*[]Gateway, error) {
	var err error
	datas := []Gateway{}
	err = db.Model(&Gateway{}).Preload("Project").Preload("User").Order("id desc").Find(&datas).Error
	if err != nil {
		return &[]Gateway{}, err
	}
	return &datas, nil
}

func (data *Gateway) Find(db *gorm.DB, pid uint64) (*Gateway, error) {
	var err error
	err = db.Model(&Gateway{}).Preload("Project").Where("id = ?", pid).Take(&data).Error
	if err != nil {
		return &Gateway{}, err
	}
	return data, nil
}