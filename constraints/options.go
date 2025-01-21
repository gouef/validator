package constraints

import (
	"errors"
	"fmt"
)

type Options struct {
	Options []any
}

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
