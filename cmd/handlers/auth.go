package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mcbryan1/resume-builder-backend/cmd/helpers"
	"github.com/mcbryan1/resume-builder-backend/cmd/models"
	"github.com/mcbryan1/resume-builder-backend/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(c *gin.Context) {
	_, user, tokenString, err := helpers.ProcessLogin(c)

	if err != nil {
		// helpers.RespondWithError(c, 400, "Login failed", err.Error())
		return
	}

	userResponse := helpers.LoginResponseSerializer(user)
	helpers.RespondWithSuccess(c, 200, "Login successful", helpers.SuccessRespCode, gin.H{
		"user":  userResponse,
		"token": tokenString,
	})
}

func RegisterUser(c *gin.Context) {
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, "Invalid request body", "400")
		return
	}

	// Check if the required fields are present in the request
	if err := helpers.ValidateRequest(req, "User"); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, err.Error(), "001")
		return
	}

	// Check if the user already exists
	email := req["email"].(string)
	if helpers.UserExistsByEmail(email) {
		helpers.RespondWithError(c, http.StatusConflict, "User already exists", "001")
		return
	}

	// Hash the password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req["password"].(string)), bcrypt.DefaultCost)
	if err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to hash password", "500")
		return
	}

	// Create a new user
	newUser := models.User{
		FirstName: req["first_name"].(string),
		LastName:  req["last_name"].(string),
		Email:     email,
		Password:  string(hashPassword),
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to create user", "500")
		return
	}

	helpers.RespondWithSuccess(c, http.StatusCreated, "User created successfully", helpers.SuccessRespCode, nil)

}

func GetUserProfile(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get user profile",
	})
}
