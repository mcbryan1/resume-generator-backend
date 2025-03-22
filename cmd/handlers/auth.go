package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mcbryan1/resume-builder-backend/cmd/helpers"
)

func LoginUser(c *gin.Context) {
	_, user, tokenString, err := helpers.ProcessLogin(c)

	if err != nil {
		// helpers.RespondWithError(c, 400, "Login failed", err.Error())
		return
	}

	userResponse := helpers.LoginResponseSerializer(user)
	helpers.RespondWithSuccess(c, 200, "Login successful", "200", gin.H{
		"user":  userResponse,
		"token": tokenString,
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
