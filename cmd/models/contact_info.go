package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ContactInfo struct {
	gorm.Model `json:"-"`
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ResumeID   uuid.UUID `json:"-" gorm:"type:uuid;not null;index"`
	FullName   string    `json:"full_name" gorm:"type:varchar(255);not null"`
	Email      string    `json:"email" gorm:"type:varchar(255);not null"`
	Phone      string    `json:"phone,omitempty" gorm:"type:varchar(50)"`
	Location   string    `json:"location,omitempty" gorm:"type:varchar(255)"`
	LinkedIn   string    `json:"linkedin,omitempty" gorm:"type:varchar(255)"`
	GitHub     string    `json:"github,omitempty" gorm:"type:varchar(255)"`
	Website    string    `json:"website,omitempty" gorm:"type:varchar(255)"`
	Others     string    `json:"others,omitempty" gorm:"type:text"`
}
