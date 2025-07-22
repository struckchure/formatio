package app

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type StructValidator struct {
	validate *validator.Validate
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationFieldErrors []ValidationError

// Error implements error.
func (v ValidationFieldErrors) Error() string {
	return "validation error"
}

// Validator needs to implement the Validate method
func (v *StructValidator) Validate(out any) error {
	// TODO: not sure if this work as intended
	err := v.validate.Struct(out)
	if err != nil {
		fieldErrors := make(ValidationFieldErrors, 0)

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		for _, err := range err.(validator.ValidationErrors) {
			fieldErrors = append(fieldErrors, ValidationError{
				Field:   err.Field(),
				Message: err.Error(),
			})
		}

		return fieldErrors
	}

	return err
}

func NewStructValidator() *StructValidator {
	validate := validator.New(validator.WithRequiredStructEnabled())

	validate.RegisterValidation("optional", func(fl validator.FieldLevel) bool { return true })
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return &StructValidator{validate: validate}
}
