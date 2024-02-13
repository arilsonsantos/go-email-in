package internalerrors

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"strings"
)

func ValidateStruct(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)

	if err == nil {
		return nil
	}

	var validationErrors validator.ValidationErrors
	errors.As(err, &validationErrors)
	validationError := validationErrors[0]

	field := strings.ToLower(validationError.StructField())
	switch validationError.Tag() {
	case "required":
		return errors.New(field + " is required")
	case "email":
		return errors.New(field + " is not a valid email")
	case "min":
		return errors.New(field + " is less than the minimum " + validationError.Param())
	case "max":
		return errors.New(field + " is greater than the maximum " + validationError.Param())
	}

	return nil
}
