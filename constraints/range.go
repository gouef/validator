package constraints

import (
	"errors"
	"fmt"
)

type RangeConstraint struct {
	Min int
	Max int
}

func (r RangeConstraint) Validate(value any) error {
	num, ok := value.(int)
	if !ok {
		return errors.New("this value must be an integer")
	}
	if num < r.Min || num > r.Max {
		return fmt.Errorf("this value must be between %d and %d", r.Min, r.Max)
	}
	return nil
}
