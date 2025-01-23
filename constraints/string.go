package constraints

import "errors"

type IsString struct{}

// Validate function for validate value
//
// Example:
//
//	con := constraints.IsString{}
//	errs := validator.Validate(value, con)
func (c IsString) Validate(value any) error {
	if _, ok := value.(string); ok {
		return nil
	}

	return errors.New("this value should be string type")
}
