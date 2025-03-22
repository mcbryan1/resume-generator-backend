package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"` // Exclude timestamps from JSON responses
	ID         uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	FirstName  string     `json:"first_name" gorm:"type:varchar(255);not null"`
	LastName   string     `json:"last_name" gorm:"type:varchar(255);not null"`
	IsActive   bool       `json:"is_active" gorm:"type:boolean;not null;default:true"`
	Email      string     `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password   string     `json:"-" gorm:"type:varchar(255);not null"` // Hide password from JSON responses
	Resumes    []Resume   `gorm:"foreignKey:UserID"`
}
