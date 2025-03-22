package main

import (
	"os"

	"github.com/mcbryan1/resume-builder-backend/internal/database"
	"github.com/mcbryan1/resume-builder-backend/internal/server"
)

func init() {
	database.LoadEnvVariables()
	database.ConnectDatabase()
}

func main() {
	r := server.RegisterRoutes()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
