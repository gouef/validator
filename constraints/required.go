package constraints

import "errors"

type Required struct{}

// Validate function for validate value
//
// Example:
//
//	con := constraints.Required{}
//	errs := validator.Validate(value, con)
func (c Required) Validate(value any) error {
	if value == nil || value == "" {
		return errors.New("this value is required")
	}

	return nil
}
