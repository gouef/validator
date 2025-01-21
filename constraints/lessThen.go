package constraints

import (
	"errors"
	"fmt"
)

type LessThen struct {
	Value float64
}

func (r LessThen) Validate(value any) error {
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

	if num >= r.Value {
		return fmt.Errorf("this value (%.2f) must be less then %.2f", num, r.Value)
	}

	return nil
}
