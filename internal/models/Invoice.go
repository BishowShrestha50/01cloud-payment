package models

import (
	"time"

	"github.com/Azure/go-autorest/autorest/date"
	"github.com/jinzhu/gorm"
)

type Invoice struct {
	gorm.Model
	ID          uint       `gorm:"foreignkey:UserID" json:"user"`
	UserID      uint       `gorm:"not null" json:"user_id"`
	Date        date.Date  `gorm:"not null" json:"date"`
	StartDate   time.Time  `gorm:"not null" json:"start_date"`
	EndDate     time.Time  `gorm:"not null" json:"end_date"`
	TotalCost   uint       `gorm:"not null" json:"total_cost"`
	PromoCode   *PromoCode `gorm:"foreignkey:PromoCodeID" json:"promocode"`
	PromoCodeID uint       `gorm:"null" json:"promo_code_id"`
	Deduction   *Deduction `gorm:"foreignkey:DeductionID" json:"deduction"`
	DeductionID uint       `gorm:"null" json:"deduction_id"`
}
func (data *Invoice) Save(db *gorm.DB) (*Invoice, error) {
	var err error
	err = db.Model(&Invoice{}).Create(&data).Error
	if err != nil {
		return &Invoice{}, err
	}
	return data, nil
}

func (data *Invoice) FindAll(db *gorm.DB) (*[]Invoice, error) {
	var err error
	datas := []Invoice{}
	err = db.Model(&Invoice{}).Preload("Project").Preload("User").Order("id desc").Find(&datas).Error
	if err != nil {
		return &[]Invoice{}, err
	}
	return &datas, nil
}

func (data *Invoice) Find(db *gorm.DB, pid uint64) (*Invoice, error) {
	var err error
	err = db.Model(&Invoice{}).Preload("Project").Where("id = ?", pid).Take(&data).Error
	if err != nil {
		return &Invoice{}, err
	}
	return data, nil
}