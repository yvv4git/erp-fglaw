package forms

import "github.com/asaskevich/govalidator"

// Base ...
type Base struct {
	errors []error
}

// GetErrors ...
func (f *Base) GetErrors() []error {
	return f.errors
}

// Validate ...
// func (f *Base) Validate() bool {
// 	return true
// }

// Validate ...
func (f *Base) Validate(form interface{}) bool {
	result, err := govalidator.ValidateStruct(form)
	if err != nil {
		f.errors = append(f.errors, err)
	}
	return result
}
