package constraints

import (
	"errors"
	"fmt"
)

type EqualTo struct {
	Value any
}

// Validate function for validate value
//
// Example:
//
//	con := constraints.EqualTo{
//		Value: "master"
//	}
//	errs := validator.Validate(value, con)
func (r EqualTo) Validate(value any) error {
	if value != r.Value {
		return errors.New(fmt.Sprintf("this value must be %v", r.Value))
	}

	return nil
}
