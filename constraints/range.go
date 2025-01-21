package constraints

import (
	"errors"
	"fmt"
)

type RangeConstraint struct {
	Min float64
	Max float64
}

func (r RangeConstraint) Validate(value any) error {
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

	if num < r.Min || num > r.Max {
		return fmt.Errorf("this value must be between %.2f and %.2f", r.Min, r.Max)
	}

	return nil
}
