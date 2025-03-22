package helpers

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mcbryan1/resume-builder-backend/cmd/models"
	"github.com/mcbryan1/resume-builder-backend/internal/database"
	"golang.org/x/crypto/bcrypt"
)

// ParseRequest parses the JSON body of a request into a map and returns it.
// It takes a *gin.Context as an argument and returns a map[string]interface{} and an error.
// If the JSON binding fails, the error will be non-nil.
// ParseRequest parses the JSON request body into a map and returns it.
//
// @param c *gin.Context - the Gin context containing the request.
// @returns map[string]interface{} - the parsed request body as a map.
// @returns error - an error if the JSON binding fails, otherwise nil.
func ParseRequest(c *gin.Context) (map[string]interface{}, error) {
	var req map[string]interface{}
	err := c.ShouldBindJSON(&req)
	return req, err
}

// CheckPassword compares the hashed password stored in the user model with the provided plain text password.
// It returns an error if the passwords do not match or if there is an issue with the comparison process.
//
// @Params:
//   - user (models.User): The user model containing the hashed password.
//   - password (string): The plain text password to compare.
//
// @Returns:
//   - error: An error if the passwords do not match or if there is an issue with the comparison process.
func CheckPassword(user models.User, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

// GenerateJWTToken generates a JWT token for the given user.
// The token is signed using the secret key stored in the environment variable "JWT_SECRET".
// The token contains the user's ID and an expiration time set to 24 hours from the time of generation.
//
// Parameters:
//   - user: The user for whom the JWT token is being generated.
//
// Returns:
//   - string: The signed JWT token.
//   - error: An error if the token could not be signed.
func GenerateJWTToken(user models.User) (string, error) {
	var s = os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(s))
}

// GetUser retrieves a user from the database based on the provided email.
//
// @param:
//   - email: A string representing the email of the user to be retrieved.
//
// @returns:
//   - user: A models.User struct representing the user retrieved from the database.
//   - error: An error if the user could not be retrieved.
func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

// ProcessLogin is a helper function that processes the login request.
// It parses the request, retrieves the user from the database, checks the password, and generates a JWT token.
// If any of these steps fail, an error response is returned.
//
// @param c *gin.Context - the Gin context containing the request.
// @returns map[string]interface{} - the parsed request body as a map.
// @returns models.User - the user retrieved from the database.
// @returns string - the generated JWT token.
// @returns error - an error if any of the steps fail.
func ProcessLogin(c *gin.Context) (req map[string]interface{}, user models.User, tokenString string, err error) {
	req, err = ParseRequest(c)
	if err != nil {
		RespondWithError(c, http.StatusBadRequest, InvalidRequestDataRespDesc, InvalidInputRespCode)
	}

	user, err = GetUserByEmail(req["email"].(string))
	if err != nil {
		RespondWithError(c, http.StatusUnauthorized, InvalidPhoneOrPasswordRespDesc, ErrorRespCode)
		return
	}

	if err = CheckPassword(user, req["password"].(string)); err != nil {
		RespondWithError(c, http.StatusUnauthorized, InvalidPhoneOrPasswordRespDesc, ErrorRespCode)
		return
	}

	tokenString, err = GenerateJWTToken(user)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, InternalServerErrorRespDesc, InternalServerErrorRespCode)
		return
	}

	return
}
