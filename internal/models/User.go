package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName     string `gorm:"size:255;not null;" json:"first_name,omitempty"`
	LastName      string `gorm:"size:255;not null;" json:"last_name,omitempty"`
	Email         string `gorm:"size:255;not null;unique" json:"email,omitempty"`
	Image         string `gorm:"size:255;null;" json:"image,omitempty"`
	Company       string `gorm:"size:255;null;" json:"company,omitempty"`
	Designation   string `gorm:"size:255;null;" json:"designation,omitempty"`
	Password      string `gorm:"size:100;not null;" json:"password,omitempty"`
	EmailVerified bool   `gorm:"not null;default:false" json:"email_verified"`
	Active        bool   `gorm:"not null;default:true" json:"active"`
	IsAdmin       bool   `gorm:"not null;default:false" json:"is_admin,omitempty"`
}

func (u *User) FindUserByID(db *gorm.DB, uid uint) (*User, error) {
	var err error
	err = db.Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}
