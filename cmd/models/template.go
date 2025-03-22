package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Template struct {
	gorm.Model `json:"-"`
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name       string    `gorm:"not null;unique"`
	PreviewURL string    `gorm:"not null"`           // URL to preview image
	Layout     string    `gorm:"type:text;not null"` // JSON structure for frontend rendering
	IsPremium  bool      `gorm:"default:false"`      // If it's a premium template
	Price      float64   `gorm:"default:0.0"`        // Price for premium templates
}
