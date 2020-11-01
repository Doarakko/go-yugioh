package yugioh

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

var validate validator.Validate

func init() {
	validate = *validator.New()
	validate.RegisterValidation("datetime", datetimeValidation)
}

// Validate ...
func Validate(input interface{}) (err error) {
	if err = validate.Struct(input); err != nil {
		return
	}
	return
}

func datetimeValidation(fl validator.FieldLevel) bool {
	_, err := time.Parse("2006-01-02 15:04:05", fl.Field().String())
	if err != nil {
		return false
	}
	return true
}
