package models

import (
	"github.com/jinzhu/gorm"
)

type PaymentHistory struct {
	gorm.Model
	User          *User        `gorm:"foreignkey:UserID" json:"user"`
	UserID        uint         `gorm:"not null" json:"user_id"`
	Debit         uint         `gorm:"not null" json:"total_cost"`
	Credit        uint         `gorm:"not null" json:"promo_code_id"`
	Balance       uint         `gorm:"not null" json:"balance"`
	Invoice       *Invoice     `gorm:"foreignkey:InvoiceID" json:"invoice"`
	InvoiceID     uint         `gorm:"not null" json:"invoice_id"`
	Transaction   *Transaction `gorm:"foreignkey:TransactionID" json:"transaction"`
	TransactionID uint         `gorm:"not null" json:"transaction_id"`
}

func (data *PaymentHistory) Save(db *gorm.DB) (*PaymentHistory, error) {
	var err error
	err = db.Model(&PaymentHistory{}).Create(&data).Error
	if err != nil {
		return &PaymentHistory{}, err
	}
	return data, nil
}

func (data *PaymentHistory) FindAll(db *gorm.DB) (*[]PaymentHistory, error) {
	var err error
	datas := []PaymentHistory{}
	err = db.Model(&PaymentHistory{}).Preload("Invoice").Preload("Transaction").Order("id desc").Find(&datas).Error
	if err != nil {
		return &[]PaymentHistory{}, err
	}
	return &datas, nil
}

func (data *PaymentHistory) Find(db *gorm.DB, pid uint64) (*PaymentHistory, error) {
	var err error
	err = db.Model(&PaymentHistory{}).Preload("Invoice").Where("id = ?", pid).Take(&data).Error
	if err != nil {
		return &PaymentHistory{}, err
	}
	return data, nil
}
