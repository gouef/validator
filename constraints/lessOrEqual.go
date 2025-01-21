package constraints

import (
	"errors"
	"fmt"
)

type LessOrEqual struct {
	Value float64
}

func (r LessOrEqual) Validate(value any) error {
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

	if num > r.Value {
		return fmt.Errorf("this value must be less or equal then %.2f", r.Value)
	}

	return nil
}
