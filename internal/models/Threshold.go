package models

import (
	"github.com/jinzhu/gorm"
)

type Threshold struct {
	gorm.Model
	User           uint   `gorm:"foreignkey:UserID" json:"user"`
	UserID         uint   `gorm:"not null" json:"user_id"`
	ThresholdLimit uint   `gorm:"not null" json:"threshold_limit"`
	Email          string `gorm:"not null" json:"email"`
	Active         bool   `gorm:"not null" json:"active"`
}
func (data *Threshold) Save(db *gorm.DB) (*Threshold, error) {
	var err error
	err = db.Model(&Threshold{}).Create(&data).Error
	if err != nil {
		return &Threshold{}, err
	}
	return data, nil
}

func (data *Threshold) FindAll(db *gorm.DB) (*[]Threshold, error) {
	var err error
	datas := []Threshold{}
	err = db.Model(&Threshold{}).Order("id desc").Find(&datas).Error
	if err != nil {
		return &[]Threshold{}, err
	}
	return &datas, nil
}

func (data *Threshold) Find(db *gorm.DB, pid uint64) (*Threshold, error) {
	var err error
	err = db.Model(&Threshold{}).Where("id = ?", pid).Take(&data).Error
	if err != nil {
		return &Threshold{}, err
	}
	return data, nil
}
