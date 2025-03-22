package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Education struct {
	gorm.Model  `json:"-"`
	ID          uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ResumeID    uuid.UUID  `gorm:"type:uuid;not null;index"`
	Institution string     `gorm:"type:varchar(255);not null"`
	Degree      string     `gorm:"type:varchar(255);not null"`
	Field       string     `gorm:"type:varchar(255)"`
	StartDate   time.Time  `gorm:"type:date;not null"`
	EndDate     *time.Time `gorm:"type:date"`
}
