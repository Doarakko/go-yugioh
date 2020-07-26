package yugioh

import "gopkg.in/go-playground/validator.v9"

var validate validator.Validate

func init() {
	validate = *validator.New()
}

// Validate ...
func Validate(input interface{}) (err error) {
	if err = validate.Struct(input); err != nil {
		return
	}
	return
}
