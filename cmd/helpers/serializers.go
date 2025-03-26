package helpers

import "github.com/mcbryan1/resume-builder-backend/cmd/models"

func LoginResponseSerializer(user models.User) map[string]interface{} {
	return map[string]interface{}{
		"id":         user.ID,
		"email":      user.Email,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"is_active":  user.IsActive,
		// "role":                RoleResponseSerializer(user.Role),
	}
}

func TemplateResponseSerializer(template models.Template) map[string]interface{} {
	return map[string]interface{}{
		"id":          template.ID,
		"name":        template.Name,
		"preview_url": template.PreviewURL,
		"is_premium":  template.IsPremium,
		"price":       template.Price,
	}
}
