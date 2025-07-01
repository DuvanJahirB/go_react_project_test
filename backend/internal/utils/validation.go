package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) string {
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return "Invalid validation error"
	}

	for _, fieldErr := range validationErrors {
		field := strings.ToLower(fieldErr.Field())
		switch fieldErr.Tag() {
		case "required":
			return fmt.Sprintf("%s is required", field)
		case "email":
			return fmt.Sprintf("%s is not a valid email", field)
		case "min":
			return fmt.Sprintf("%s must be at least %s characters long", field, fieldErr.Param())
		default:
			return fmt.Sprintf("%s is not valid", field)
		}
	}

	return "Validation error"
}