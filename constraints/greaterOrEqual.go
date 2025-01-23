package constraints

import (
	"errors"
	"fmt"
)

type GreaterOrEqual struct {
	Value float64
}

// Validate function for validate value
//
// Example:
//
//	con := constraints.GreaterOrEqual{
//		Value: 35
//	}
//	errs := validator.Validate(value, con)
func (r GreaterOrEqual) Validate(value any) error {
	var num float64

	switch v := value.(type) {
	case int:
		num = float64(v)
	case float32:
		num = float64(v)
	case float64:
		num = v
	default:
		return errors.New("this value must be an integer or float")
	}

	if num < r.Value {
		return fmt.Errorf("this value must be greater or equal then %.2f", r.Value)
	}

	return nil
}
