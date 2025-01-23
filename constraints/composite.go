package constraints

import (
	"fmt"
	"github.com/gouef/validator"
)

type Composite struct {
	Constraints []validator.Constraint
}

// Validate function for validate value
//
// Example:
//
//		nestedValidator := constraints.Field{
//		   Constraints: map[string]validator.Constraint{
//		       "Email": constraints.Composite{
//		           Constraints: []validator.Constraint{
//		               constraints.NotBlank{},
//		               constraints.Length{Min: 5, Max: 50},
//		               constraints.Email{},
//		           },
//		       },
//		       "Name": constraints.Composite{
//		           Constraints: []validator.Constraint{
//		               constraints.NotBlank{},
//		               constraints.Length{Min: 3, Max: 100},
//		           },
//		       },
//		   },
//		}
//
//	 errs := validator.Validate(value, nestedValidator)
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
