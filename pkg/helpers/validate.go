package helpers

import (
	"github.com/go-playground/validator/v10"
)

func ValidateRequestBody(b interface{}) error {
	validate := validator.New()
	return validate.Struct(b)
}
