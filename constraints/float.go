package constraints

import "errors"

type IsFloat struct{}
type IsFloat32 struct{}
type IsFloat64 struct{}

func (c IsFloat) Validate(value any) error {
	if _, ok := value.(float64); ok {
		return nil
	}
	if _, ok2 := value.(float32); ok2 {
		return nil
	}

	return errors.New("this value should be float type")
}

func (c IsFloat32) Validate(value any) error {
	if _, ok := value.(float32); ok {
		return nil
	}

	return errors.New("this value should be float32 type")
}

func (c IsFloat64) Validate(value any) error {
	if _, ok := value.(float64); ok {
		return nil
	}

	return errors.New("this value should be float64 type")
}
