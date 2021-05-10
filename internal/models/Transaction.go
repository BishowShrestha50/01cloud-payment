package models

import (
	"github.com/jinzhu/gorm"
)

type Transaction struct {
	gorm.Model
	Title     string   `gorm:"not null" json:"title"`
	User      *User    `gorm:"foreignkey:UserID" json:"user"`
	UserID    uint     `gorm:"not null" json:"user_id"`
	Amount    uint     `gorm:"not null" json:"amount"`
	Credit    uint     `gorm:"not null" json:"promo_code_id"`
	Balance   uint     `gorm:"not null" json:"balance"`
	Invoice   *Invoice `gorm:"foreignkey:InvoiceID" json:"invoice"`
	InvoiceID uint     `gorm:"not null" json:"invoice_id"`
	Gateway   *Gateway `gorm:"foreignkey:GatewayID" json:"gateway"`
	GatewayID uint     `gorm:"not null" json:"gateway_id"`
	Status    string   `gorm:"not null" json:"status"`
}
func (data *Transaction) Save(db *gorm.DB) (*Transaction, error) {
	var err error
	err = db.Model(&Transaction{}).Create(&data).Error
	if err != nil {
		return &Transaction{}, err
	}
	return data, nil
}

func (data *Transaction) FindAll(db *gorm.DB) (*[]Transaction, error) {
	var err error
	datas := []Transaction{}
	err = db.Model(&Transaction{}).Order("id desc").Find(&datas).Error
	if err != nil {
		return &[]Transaction{}, err
	}
	return &datas, nil
}

func (data *Transaction) Find(db *gorm.DB, pid uint64) (*Transaction, error) {
	var err error
	err = db.Model(&Transaction{}).Where("id = ?", pid).Take(&data).Error
	if err != nil {
		return &Transaction{}, err
	}
	return data, nil
}
