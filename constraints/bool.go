package constraints

import "errors"

type IsFalse struct{}
type IsTrue struct{}
type IsBool struct{}

// Validate function for validate value
//
// Example:
//
//	con := constraints.IsFalse{}
//	errs := validator.Validate(value, con)
func (c IsFalse) Validate(value any) error {
	if value == false {
		return nil
	}

	return errors.New("this value should be false")
}

// Validate function for validate value
//
// Example:
//
//	con := constraints.IsTrue{}
//	errs := validator.Validate(value, con)
func (c IsTrue) Validate(value any) error {
	if value == true {
		return nil
	}

	return errors.New("this value should not be true")
}

// Validate function for validate value
//
// Example:
//
//	con := constraints.IsBool{}
//	errs := validator.Validate(value, con)
func (c IsBool) Validate(value any) error {
	if _, ok := value.(bool); ok {
		return nil
	}

	return errors.New("this value should not be true")
}
