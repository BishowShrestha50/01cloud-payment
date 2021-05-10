package models

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type Subscription struct {
	gorm.Model
	Name         string         `gorm:"size:255;not null;" json:"name"`
	Apps         uint32         `gorm:"not null;default:1" json:"apps"`
	DiskSpace    uint32         `gorm:"not null;" json:"disk_space"`
	Memory       uint32         `gorm:"not null;" json:"memory"`
	Cores        uint32         `gorm:"not null;" json:"cores"`
	DataTransfer uint32         `gorm:"not null;" json:"data_transfer"`
	Price        uint32         `gorm:"not null;" json:"price"`
	Weight       uint32         `gorm:"default:10;" json:"weight"`
	Attributes   string         `gorm:"size:1024;null;" json:"attributes"`
	Active       bool           `gorm:"not null;" json:"active"`
	CronJob      uint64         `gorm:"default:1;" json:"cron_job"`
	Backups      uint64         `gorm:"default:5;" json:"backups"`
	ResourceList postgres.Jsonb `sql:"json" json:"resource_list"`
	//Organization   *Organization  `gorm:"foreignkey:OrganizationID" json:"organization,omitempty"`
	//OrganizationID uint64         `gorm:"default:0" json:"organization_id"`
	LoadBalancer uint64 `gorm:"default:0;" json:"load_balancer"`
}

func (data *Subscription) Find(db *gorm.DB, pid uint64) (*Subscription, error) {
	var err error
	err = db.Model(&Subscription{}).
		Where("id = ?", pid).
		Take(&data).Error
	if err != nil {
		return &Subscription{}, err
	}
	return data, nil
}
