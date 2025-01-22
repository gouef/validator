package constraints

import (
	"errors"
	"fmt"
	"reflect"
)

type Length struct {
	Min int
	Max int
}

func (c Length) Validate(value any) error {
	if value == nil {
		return errors.New("this value is required and cannot be nil")
	}

	v := reflect.ValueOf(value)
	length := 0

	switch v.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
		length = v.Len()
	default:
		return errors.New("this value must be a string, array, slice, or map")
	}

	if c.Min > 0 && length < c.Min {
		return fmt.Errorf("this value must have a length of at least %d", c.Min)
	}
	if c.Max > 0 && length > c.Max {
		return fmt.Errorf("this value must have a length of at most %d", c.Max)
	}

	return nil
}
