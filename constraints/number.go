package constraints

import (
	"errors"
	"fmt"
)

type IsEven struct{}

func (c IsEven) Validate(value any) error {
	v, ok := value.(int)
	if !ok {
		return errors.New("this value should be int type")
	}
	if v%2 != 0 {
		return errors.New("this value should be even")
	}
	return nil
}

type IsOdd struct{}

func (c IsOdd) Validate(value any) error {
	v, ok := value.(int)
	if !ok {
		return errors.New("this value should be int type")
	}
	if v%2 == 0 {
		return errors.New("this value should be odd")
	}
	return nil
}

type IsMultipleOf struct {
	Divisor int
}

func (c IsMultipleOf) Validate(value any) error {
	v, ok := value.(int)
	if !ok {
		return errors.New("this value should be int type")
	}
	if c.Divisor == 0 {
		return errors.New("divisor cannot be zero")
	}
	if v%c.Divisor != 0 {
		return fmt.Errorf("this value should be a multiple of %d", c.Divisor)
	}
	return nil
}
