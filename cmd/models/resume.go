package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Resume struct {
	gorm.Model     `json:"-"`
	ID             uuid.UUID        `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID         uuid.UUID        `json:"-" gorm:"type:uuid;not null;index"` // Hide foreign key
	User           User             `json:"-" gorm:"foreignKey:UserID"`
	TemplateID     *uuid.UUID       `json:"template_id,omitempty" gorm:"type:uuid;default:null"`
	Template       *Template        `json:"template,omitempty" gorm:"foreignKey:TemplateID"`
	Title          string           `json:"title" gorm:"type:varchar(255);not null"`
	ContactInfo    ContactInfo      `json:"contact_info" gorm:"foreignKey:ResumeID;constraint:OnDelete:CASCADE;"`
	Summary        string           `json:"summary" gorm:"type:text"`
	Skills         []*Skill         `json:"skills,omitempty" gorm:"many2many:resume_skills;"`
	Experience     []WorkExperience `json:"experience,omitempty" gorm:"foreignKey:ResumeID;constraint:OnDelete:CASCADE;"`
	Projects       []Project        `json:"projects,omitempty" gorm:"foreignKey:ResumeID;constraint:OnDelete:CASCADE;"`
	Education      []Education      `json:"education,omitempty" gorm:"foreignKey:ResumeID;constraint:OnDelete:CASCADE;"`
	Certifications []Certification  `json:"certifications,omitempty" gorm:"foreignKey:ResumeID;constraint:OnDelete:CASCADE;"`
}
