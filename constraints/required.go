package constraints

import "errors"

type Required struct{}

func (c Required) Validate(value any) error {
	if value == nil || value == "" {
		return errors.New("this value is required")
	}

	return nil
}
