package constraints

import (
	"errors"
	"fmt"
)

type ArrayRangeConstraint struct {
	Min float64
	Max float64
}

func (c ArrayRangeConstraint) Validate(value any) error {
	slice, ok := value.([]any)
	if !ok {
		return errors.New("this value must be an array or slice")
	}

	for i, item := range slice {
		var num float64

		switch v := item.(type) {
		case int:
			num = float64(v)
		case float32:
			num = float64(v)
		case float64:
			num = v
		default:
			return fmt.Errorf("element at index %d is not a number", i)
		}

		if num < c.Min || num > c.Max {
			return fmt.Errorf("element at index %d must be between %.2f and %.2f, got %.2f", i, c.Min, c.Max, num)
		}
	}

	return nil
}
