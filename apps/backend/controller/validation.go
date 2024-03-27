package controller

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

// Converts validation errors to a map
func validationErrorsToMap(valErrs validator.ValidationErrors) map[string]string {
	errors := make(map[string]string)
	for _, valErr := range valErrs {
		// Field() gives the field name, Tag() gives the validation rule that failed
		errors[valErr.Field()] = fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", valErr.Field(), valErr.Tag())
	}
	return errors
}
