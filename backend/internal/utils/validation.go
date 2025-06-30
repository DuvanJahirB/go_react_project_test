package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) map[string]string {
	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		field := strings.ToLower(err.Field())
		switch err.Tag() {
		case "required":
			errors[field] = fmt.Sprintf("%s is required", field)
		case "email":
			errors[field] = fmt.Sprintf("%s is not a valid email", field)
		case "min":
			errors[field] = fmt.Sprintf("%s must be at least %s characters long", field, err.Param())
		default:
			errors[field] = fmt.Sprintf("%s is not valid", field)
		}
	}
	return errors
}