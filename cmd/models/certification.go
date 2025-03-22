package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Certification struct {
	gorm.Model   `json:"-"`
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ResumeID     uuid.UUID `gorm:"type:uuid;not null;index"`
	Title        string    `gorm:"type:varchar(255);not null"`
	Issuer       string    `gorm:"type:varchar(255)"`
	DateReceived time.Time `gorm:"type:date"`
}
