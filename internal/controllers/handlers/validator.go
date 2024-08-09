package handlers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

var (
	categories = [...]string{"music", "funny", "videos", "programming", "news", "fashion"}
)

func CategoryValidationFunc(fl validator.FieldLevel) bool {
	input := fl.Field().String()
	for _, category := range categories {
		if input == category {
			return true
		}
	}
	return false
}

func NewValidator() (*validator.Validate, error) {
	v := validator.New()

	err := v.RegisterValidation("categoryValidation", CategoryValidationFunc)
	if err != nil {
		return nil, fmt.Errorf("couldn't register validator: %w", err)
	}

	return v, nil
}

func ValidationError(errs validator.ValidationErrors) string {
	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is a required field", err.Field()))
		case "url":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not a valid URL", err.Field()))
		case "alphanum":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is neither num nor alpha", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not valid", err.Field()))
		}
	}

	return strings.Join(errMsgs, ", ")
}
