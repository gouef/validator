package constraints

import (
"errors"
"fmt"
)

type EqualTo struct {
	Value any
}

func (r EqualTo) Validate(value any) error {
	if value != r.Value {
		return errors.New(fmt.Sprintf("this value must be %v", r.Value))
	}

	return nil
}
