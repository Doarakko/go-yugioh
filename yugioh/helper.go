package yugioh

import (
	"log"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate validator.Validate

func init() {
	validate = *validator.New()
	err := validate.RegisterValidation("datetime", datetimeValidation)
	if err != nil {
		log.Fatal("failed to register validation")
	}
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
	return err == nil
}
