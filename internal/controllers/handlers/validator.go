package handlers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

var (
	categories = [...]string{"music", "funny", "videos", "programming", "news", "fashion"}
	postTypes  = [...]string{"text", "link"}
)

type Validator struct {
	*validator.Validate
}

func (v *Validator) AddPostValidator(input *AddPostInput) error {
	switch input.PostType {
	case "url":
		if input.URL == "" {
			return fmt.Errorf("url field should not be empty for posts with url type")
		}
	case "text":
		if input.Text == "" {
			return fmt.Errorf("text field should not be empty for posts with text type")
		}
	}

	return nil
}

func CategoryValidationFunc(fl validator.FieldLevel) bool {
	input := fl.Field().String()
	for _, category := range categories {
		if input == category {
			return true
		}
	}
	return false
}

func PostTypeValidationFunc(fl validator.FieldLevel) bool {
	input := fl.Field().String()
	for _, postType := range postTypes {
		if input == postType {
			return true
		}
	}
	return false
}

func NewValidator() (*Validator, error) {
	v := &Validator{
		validator.New(),
	}

	err := v.RegisterValidation("categoryValidation", CategoryValidationFunc)
	if err != nil {
		return nil, fmt.Errorf("couldn't register categoryValidation validator: %w", err)
	}

	err = v.RegisterValidation("postTypeValidation", PostTypeValidationFunc)
	if err != nil {
		return nil, fmt.Errorf("couldn't register postTypeValidation validator: %w", err)
	}

	return v, nil
}

func ValidationError(errs validator.ValidationErrors) string {
	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "postTypeValidation":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not a valid postType", err.Field()))
		case "categoryValidation":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not a valid category", err.Field()))
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
