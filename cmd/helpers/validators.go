package helpers

import (
	"fmt"
	"strings"
)

func ValidateRequest(req map[string]interface{}, req_type string) error {
	var requiredFields []string

	switch req_type {
	case "User":
		requiredFields = []string{"first_name", "last_name", "email", "password"}
	case "Todo":
		requiredFields = []string{"title"}
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

	return nil
}
