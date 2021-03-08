package forms

import "github.com/asaskevich/govalidator"

// Validate is used for validation all form in one point.
func Validate(form interface{}) (result bool, err error) {
	return govalidator.ValidateStruct(form)
}
