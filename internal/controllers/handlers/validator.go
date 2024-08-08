package handlers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
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
