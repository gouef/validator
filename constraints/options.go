package constraints

import (
	"errors"
	"fmt"
)

type Options struct {
	Options []any
}

// Validate function for validate value
//
// Example:
//
//	con := constraints.Options{
//		Options: []any{"option1", "option2", "option3", 42},
//	}
//
//	value := "option1"
//	errs := validator.Validate(value, con)
func (c Options) Validate(value any) error {
	for _, option := range c.Options {
		if option == value {
			return nil
		}
	}

	return fmt.Errorf("this value must be one of the allowed options: %v", c.Options)
}

type ArrayOptions struct {
	Options []any
}

// Validate function for validate value
//
// Example:
//
//	con := constraints.Options{
//		Options: []any{"Red", "Green", "Blue", 100},
//	}
//
//	values := []any{"Red", "Green", 100}
//	errs := validator.Validate(values, con)
func (c ArrayOptions) Validate(value any) error {
	arr, ok := value.([]any)
	if !ok {
		return errors.New("this value must be an array")
	}

	for i, elem := range arr {
		found := false
		for _, option := range c.Options {
			if elem == option {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("element at index %d (%v) is not one of the allowed options: %v", i, elem, c.Options)
		}
	}

	return nil
}
