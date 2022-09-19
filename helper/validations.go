package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationResponse struct {
	Success     bool         `json:"success"`
	Validations []Validation `json:"validations"`
}

type Validation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GenerateValidationMessage(field string, rule string) (message string) {
	switch rule {
	case "required":
		return fmt.Sprintf("Field '%s' is '%s'.", field, rule)
	// you can add another validator.v8 rule here
	default:
		return fmt.Sprintf("Field '%s' is not valid.", field)
	}
}

func FormatValidationError(err error) (response ValidationResponse) {
	response.Success = false

	var validations []Validation
	// get validation errors
	validationErrors := err.(validator.ValidationErrors)
	fmt.Println("err", validationErrors)

	for _, value := range validationErrors {
		// get field & rule (tag)
		field, rule := value.Field(), value.Tag()

		// create validation object
		validation := Validation{Field: field, Message: GenerateValidationMessage(field, rule)}

		// add validation object to validations
		validations = append(validations, validation)
	}

	// set Validations response
	response.Validations = validations

	return response
}
