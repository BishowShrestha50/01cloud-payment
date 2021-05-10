package models

import (
	"github.com/jinzhu/gorm"
)

type InvoiceItems struct {
	gorm.Model
	Invoice    uint   `gorm:"foreignkey:InvoiceID" json:"invoice"`
	InvoiceID  uint   `gorm:"not null" json:"invoice_id"`
	UserID     uint   `gorm:"not null" json:"user_id"`
	Particular string `gorm:"not null" json:"particular"`
	Rate       uint   `gorm:"not null" json:"rate"`
	Days       uint   `gorm:"not null" json:"days"`
	Total      uint   `gorm:"not null" json:"total"`
}

func (data *InvoiceItems) Save(db *gorm.DB) (*InvoiceItems, error) {
	var err error
	err = db.Model(&InvoiceItems{}).Create(&data).Error
	if err != nil {
		return &InvoiceItems{}, err
	}
	return data, nil
}

func (data *InvoiceItems) FindAll(db *gorm.DB) (*[]InvoiceItems, error) {
	var err error
	datas := []InvoiceItems{}
	err = db.Model(&InvoiceItems{}).Order("id desc").Find(&datas).Error
	if err != nil {
		return &[]InvoiceItems{}, err
	}
	return &datas, nil
}

func (data *InvoiceItems) Find(db *gorm.DB, pid uint64) (*InvoiceItems, error) {
	var err error
	err = db.Model(&InvoiceItems{}).Where("id = ?", pid).Take(&data).Error
	if err != nil {
		return &InvoiceItems{}, err
	}
	return data, nil
}
