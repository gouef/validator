package constraints

import "errors"

type Nil struct{}
type NotNil struct{}

func (c Nil) Validate(value any) error {
	if value == nil {
		return nil
	}

	return errors.New("this value should be nil")
}

func (c NotNil) Validate(value any) error {
	if value != nil {
		return nil
	}

	return errors.New("this value should not be nil")
}
