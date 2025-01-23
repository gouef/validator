package constraints

import "errors"

type Nil struct{}
type NotNil struct{}

// Validate function for validate value
//
// Example:
//
//	con := constraints.Nil{}
//	errs := validator.Validate(value, con)
func (c Nil) Validate(value any) error {
	if value == nil {
		return nil
	}

	return errors.New("this value should be nil")
}

// Validate function for validate value
//
// Example:
//
//	con := constraints.NotNil{}
//	errs := validator.Validate(value, con)
func (c NotNil) Validate(value any) error {
	if value != nil {
		return nil
	}

	return errors.New("this value should not be nil")
}
