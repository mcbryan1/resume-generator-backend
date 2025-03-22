package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Skill struct {
	gorm.Model `json:"-"`
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name       string    `json:"name" gorm:"type:varchar(255);unique;not null"`
	Resumes    []*Resume `json:"-" gorm:"many2many:resume_skills;"` // Hide relationship from JSON
}
