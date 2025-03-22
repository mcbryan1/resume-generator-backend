package handlers

import "github.com/gin-gonic/gin"

func LoginUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Login user",
	})
}

func RegisterUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Register user",
	})
}

func GetUserProfile(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get user profile",
	})
}
