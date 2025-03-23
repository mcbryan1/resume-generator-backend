package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mcbryan1/resume-builder-backend/cmd/handlers"
	"github.com/mcbryan1/resume-builder-backend/cmd/middlewares"
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
	auth.Use(middlewares.RateLimitMiddleware())
	{
		auth.POST("/register", handlers.RegisterUser)
		auth.POST("/login", handlers.LoginUser)
		// Protected routes group
		protected := auth.Group("")
		protected.Use(middlewares.AuthMiddleware())
		{
			protected.GET("/profile", handlers.GetUserProfile)
		}
	}
}
