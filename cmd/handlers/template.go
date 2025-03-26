package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mcbryan1/resume-builder-backend/cmd/helpers"
	"github.com/mcbryan1/resume-builder-backend/cmd/models"
	"github.com/mcbryan1/resume-builder-backend/internal/database"
)

func CreateTemplate(c *gin.Context) {
	_, ok, err := helpers.GetUserIDFromContext(c)
	if !ok || err != nil {
		helpers.RespondWithError(c, http.StatusUnauthorized, "User not authenticated", "401")
		return
	}
	var req struct {
		Name       string  `json:"name" binding:"required"`
		PreviewURL string  `json:"preview_url" binding:"required,url"`
		IsPremium  bool    `json:"is_premium"`
		Price      float64 `json:"price"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, err.Error(), "400")
		return
	}

	reqMap := map[string]interface{}{
		"name":        req.Name,
		"preview_url": req.PreviewURL,
		"is_premium":  req.IsPremium,
		"price":       req.Price,
	}

	if err := helpers.ValidateRequest(reqMap, "Template"); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, err.Error(), "400")
		return
	}
	// Avoid duplicate template names
	if helpers.TemplateExistsByName(req.Name) {
		helpers.RespondWithError(c, http.StatusConflict, "Template already exists", "409")
		return
	}

	template := models.Template{
		Name:       req.Name,
		PreviewURL: req.PreviewURL,
		IsPremium:  req.IsPremium,
		Price:      req.Price,
	}

	if err := database.DB.Create(&template).Error; err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, err.Error(), "500")
		return
	}

	helpers.RespondWithSuccess(c, http.StatusCreated, "Template created successfully", helpers.SuccessRespCode, gin.H{
		"template": helpers.TemplateResponseSerializer(template),
	})
}
