package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type Project struct {
	gorm.Model
	Name        string         `gorm:"size:255;not null;" json:"name"`
	Description string         `gorm:"size:1024;null;" json:"description"`
	ProjectCode string         `gorm:"size:5;null;" json:"project_code"`
	Tags        postgres.Jsonb `gorm:"null;" json:"tags"`
	//ClusterScope   ClusterScopeEnum `gorm:"null;" json:"cluster_scope"`   // "Shared Cluster" / "Organization Cluster"
	Region         string         `gorm:"size:127;null;" json:"region"` // Global or Regional
	Logging        postgres.Jsonb `json:"logging"`                      // blank for default Logging
	Monitoring     postgres.Jsonb `json:"monitoring"`                   // blank for default Monitoring
	BaseDomain     string         `gorm:"size:127;null;" json:"base_domain"`
	DedicatedLb    bool           `gorm:"default:false;" json:"dedicated_lb"`
	OptimizeCost   bool           `gorm:"default:false;" json:"optimize_cost"`
	Active         bool           `gorm:"not null;" json:"active"`
	Attributes     postgres.Jsonb `json:"attributes"`
	Subscription   *Subscription  `gorm:"foreignkey:SubscriptionID" json:"subscription,omitempty"`
	SubscriptionID uint64         `gorm:"not null" json:"subscription_id"`
	Image          string         `gorm:"null" json:"image,omitempty"`
	Variables      postgres.Jsonb `sql:"json" json:"variables"`
	User           *User          `gorm:"foreignkey:UserID" json:"user,omitempty"`
	UserID         uint64         `gorm:"not null" json:"user_id"`
}

func (data *Project) FindAllByUser(db *gorm.DB, userID uint, startDate, endDate time.Time) ([]Project, error) {
	dataList := []Project{}
	fmt.Println("klkl")
	db.Debug().Raw("select * from projects WHERE projects.id IN (select id from projects where user_id=?) and (projects.deleted_at BETWEEN ?::timestamp AND ?::timestamp or projects.deleted_at is null)", userID, startDate, endDate).Scan(&dataList)
	return dataList, nil
}
