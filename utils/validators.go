package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())

func ValidateInput(data interface{}) error {
	err := validate.Struct(data)

	if err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, "Invalid "+err.Field())
		}
		return fmt.Errorf(validationErrors[0])

	}
	return nil

}
