package models

import (
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/jinzhu/gorm"
)

type PromoCode struct {
	gorm.Model
	Title      string    `gorm:"not null" json:"title"`
	Code       uint      `gorm:"not null" json:"code"`
	IsPercent  bool      `gorm:"not null" json:"is_percent"`
	Discount   uint      `gorm:"not null" json:"discount"`
	ExpiryDate date.Date `gorm:"not null" json:"expiry_date"`
	Limit      uint      `gorm:"not null" json:"limit"`
	Count      uint      `gorm:"not null" json:"count"`
	Active     bool      `gorm:"not null" json:"active"`
}

func (data *PromoCode) Save(db *gorm.DB) (*PromoCode, error) {
	var err error
	err = db.Model(&PromoCode{}).Create(&data).Error
	if err != nil {
		return &PromoCode{}, err
	}
	return data, nil
}

func (data *PromoCode) FindAll(db *gorm.DB) (*[]PromoCode, error) {
	var err error
	datas := []PromoCode{}
	err = db.Model(&PromoCode{}).Order("id desc").Find(&datas).Error
	if err != nil {
		return &[]PromoCode{}, err
	}
	return &datas, nil
}

func (data *PromoCode) Find(db *gorm.DB, pid uint64) (*PromoCode, error) {
	var err error
	err = db.Model(&PromoCode{}).Where("id = ?", pid).Take(&data).Error
	if err != nil {
		return &PromoCode{}, err
	}
	return data, nil
}
