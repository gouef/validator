package constraints

import "errors"

type IsString struct{}

func (c IsString) Validate(value any) error {
	if _, ok := value.(string); ok {
		return nil
	}

	return errors.New("this value should be string type")
}
