package models

import (
	"github.com/jinzhu/gorm"
)

type Deduction struct {
	gorm.Model
	Name      string `gorm:"not null" json:"name"`
	Value     uint   `gorm:"not null" json:"value"`
	IsPercent bool   `gorm:"not null" json:"is_percent"`
	Country   string `gorm:"not null" json:"country"`
}

func (data *Deduction) Save(db *gorm.DB) (*Deduction, error) {
	var err error
	err = db.Model(&Deduction{}).Create(&data).Error
	if err != nil {
		return &Deduction{}, err
	}
	return data, nil
}

func (data *Deduction) FindAll(db *gorm.DB) (*[]Deduction, error) {
	var err error
	datas := []Deduction{}
	err = db.Model(&Deduction{}).Order("id desc").Find(&datas).Error
	if err != nil {
		return &[]Deduction{}, err
	}
	return &datas, nil
}

func (data *Deduction) Find(db *gorm.DB, pid uint64) (*Deduction, error) {
	var err error
	err = db.Model(&Deduction{}).Where("id = ?", pid).Take(&data).Error
	if err != nil {
		return &Deduction{}, err
	}
	return data, nil
}
