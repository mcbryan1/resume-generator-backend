package helpers

import (
	"encoding/json"
	"fmt"
	"strings"
)

func ValidateRequest(req map[string]interface{}, req_type string) error {
	var requiredFields []string

	switch req_type {
	case "User":
		requiredFields = []string{"first_name", "last_name", "email", "password"}
	case "Template":
		requiredFields = []string{"name", "preview_url"}
	default:
		return fmt.Errorf("invalid request type")
	}

	for _, field := range requiredFields {
		if _, ok := req[field]; !ok {
			return fmt.Errorf("%s is required", field)
		}
		// Trim whitespace from the field value if it's a string
		if strVal, ok := req[field].(string); ok {
			strVal = strings.TrimSpace(strVal)
			if strVal == "" {
				return fmt.Errorf("%s cannot be empty", field)
			}
			req[field] = strVal
		}
	}

	// Additional validation for User type
	if req_type == "User" {
		email, ok := req["email"].(string)
		if !ok || !IsEmailValid(email) {
			return fmt.Errorf("invalid email")
		}
	}

	// Additional validation for Template type
	if req_type == "Template" {
		// Validate price
		if price, ok := req["price"].(float64); ok {
			if price < 0 {
				return fmt.Errorf("price must be greater than 0")
			}
		}

		// Validate premium template price
		if isPremium, ok := req["is_premium"].(bool); ok {
			if price, ok := req["price"].(float64); ok {
				if isPremium && price == 0 {
					return fmt.Errorf("premium templates must have a price")
				}
			}
		}
	}

	return nil
}

func ValidateTemplateLayout(layout string) error {
	var jsonStruct interface{}
	if err := json.Unmarshal([]byte(layout), &jsonStruct); err != nil {
		return fmt.Errorf("invalid JSON layout: %v", err)
	}
	return nil
}
