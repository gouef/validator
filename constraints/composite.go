package constraints

import (
	"fmt"
	"github.com/gouef/validator"
)

type Composite struct {
	Constraints []validator.Constraint
}

func (c Composite) Validate(value any) error {
	var validationErrors []error

	for _, constraint := range c.Constraints {
		if err := constraint.Validate(value); err != nil {
			validationErrors = append(validationErrors, err)
		}
	}

	if len(validationErrors) > 0 {
		return fmt.Errorf("composite validation errors: %v", validationErrors)
	}
	return nil
}
