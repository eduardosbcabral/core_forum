package config

import(
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func StartValidator() {
	validate = validator.New()
}