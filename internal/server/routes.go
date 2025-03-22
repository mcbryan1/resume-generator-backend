package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mcbryan1/resume-builder-backend/cmd/handlers"
)

func RegisterRoutes() *gin.Engine {
	r := gin.Default()

	applyCORS(r)
	authRoutes(r)

	return r
}

func applyCORS(r *gin.Engine) {
	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Accept", "source"}

	r.Use(cors.New(config))
}

func authRoutes(r *gin.Engine) {
	auth := r.Group("/v1/auth")
	{
		auth.POST("/register", handlers.RegisterUser)
		auth.POST("/login", handlers.LoginUser)
		auth.GET("/profile", handlers.GetUserProfile)
	}

}
