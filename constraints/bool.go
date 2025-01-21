package constraints

import "errors"

type IsFalse struct{}
type IsTrue struct{}
type IsBool struct{}

func (c IsFalse) Validate(value any) error {
	if value == false {
		return nil
	}

	return errors.New("this value should be false")
}

func (c IsTrue) Validate(value any) error {
	if value == true {
		return nil
	}

	return errors.New("this value should not be true")
}

func (c IsBool) Validate(value any) error {
	if _, ok := value.(bool); ok {
		return nil
	}

	return errors.New("this value should not be true")
}
