package main

import (
	"log"

	"github.com/mcbryan1/resume-builder-backend/cmd/models"
	"github.com/mcbryan1/resume-builder-backend/internal/database"
)

func init() {
	database.LoadEnvVariables()
	database.ConnectDatabase()
}

func main() {

	// if err := database.DB.AutoMigrate(&models.User{}); err != nil {
	// 	log.Fatalf("Failed to migrate User model: %v", err)
	// }

	if err := database.DB.AutoMigrate(&models.Template{}); err != nil {
		log.Fatalf("Failed to migrate Template model: %v", err)
	}

	// if err := database.DB.AutoMigrate(&models.Resume{}); err != nil {
	// 	log.Fatalf("Failed to migrate Resume model: %v", err)
	// }

	// if err := database.DB.AutoMigrate(&models.ContactInfo{}); err != nil {
	// 	log.Fatalf("Failed to migrate ContactInfo model: %v", err)
	// }

	// if err := database.DB.AutoMigrate(&models.Skill{}); err != nil {
	// 	log.Fatalf("Failed to migrate Skill model: %v", err)
	// }

	// if err := database.DB.AutoMigrate(&models.WorkExperience{}); err != nil {
	// 	log.Fatalf("Failed to migrate WorkExperience model: %v", err)
	// }

	// if err := database.DB.AutoMigrate(&models.Project{}); err != nil {
	// 	log.Fatalf("Failed to migrate Project model: %v", err)
	// }

	// if err := database.DB.AutoMigrate(&models.Education{}); err != nil {
	// 	log.Fatalf("Failed to migrate Education model: %v", err)
	// }

	// if err := database.DB.AutoMigrate(&models.Certification{}); err != nil {
	// 	log.Fatalf("Failed to migrate Certification model: %v", err)
	// }

	// err := database.DB.AutoMigrate(
	// 	&models.User{},
	// 	&models.Resume{},
	// 	&models.Skill{},
	// 	&models.WorkExperience{},
	// 	&models.Project{},
	// 	&models.Education{},
	// 	&models.Certification{},
	// )

	// if err != nil {
	// 	panic("Failed to migrate the database" + err.Error())
	// } else {
	// 	println("Database migrated successfully")
	// }
}
