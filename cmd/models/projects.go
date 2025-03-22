package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model  `json:"-"`
	ID          uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ResumeID    uuid.UUID  `json:"-" gorm:"type:uuid;not null;index"`
	Title       string     `json:"title" gorm:"type:varchar(255);not null"`
	Description string     `json:"description,omitempty" gorm:"type:text"`
	Link        string     `json:"link,omitempty" gorm:"type:varchar(255)"`
	StartDate   time.Time  `json:"start_time,omitempty" gorm:"type:date;"`
	EndDate     *time.Time `json:"end_time,omitempty" gorm:"type:date"`
	ProjectType string     `json:"project_type,omitempty" gorm:"type:varchar(255)"`
}
