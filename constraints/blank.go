package constraints

import "errors"

type Blank struct{}
type NotBlank struct{}

// Validate function for validate value
//
// Example:
//
//	con := constraints.Blank{}
//	errs := validator.Validate(value, con)
func (c Blank) Validate(value any) error {
	if value == "" || value == nil {
		return nil
	}

	return errors.New("this value should be blank")
}

// Validate function for validate value
//
// Example:
//
//	con := constraints.NotBlank{}
//	errs := validator.Validate(value, con)
func (c NotBlank) Validate(value any) error {
	if value != "" && value != nil {
		return nil
	}

	return errors.New("this value should not be blank")
}
