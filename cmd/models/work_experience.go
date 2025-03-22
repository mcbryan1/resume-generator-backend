package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkExperience struct {
	gorm.Model  `json:"-"`
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ResumeID    uuid.UUID `json:"-" gorm:"type:uuid;not null;index"`
	Company     string    `json:"company" gorm:"type:varchar(255);not null"`
	Position    string    `json:"position" gorm:"type:varchar(255);not null"`
	StartDate   string    `json:"start_date" gorm:"type:varchar(20);not null"` // Use string to avoid timezone issues
	EndDate     *string   `json:"end_date,omitempty" gorm:"type:varchar(20)"`
	Description string    `json:"description,omitempty" gorm:"type:text"`
}
