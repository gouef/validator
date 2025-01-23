package constraints

import (
	"fmt"

	"github.com/gouef/validator"
)

type Field struct {
	Constraints map[string]validator.Constraint
}

// Validate function for validate value
//
// Example:
//
//	con := constraints.Field{
//	   Constraints: map[string]validator.Constraint{
//	       "Name":  constraints.NotBlank{},
//	       "Email": constraints.Email{},
//	       "Address": constraints.Field{
//	           Constraints: map[string]validator.Constraint{
//	               "City":    constraints.NotBlank{},
//	               "Country": constraints.NotBlank{},
//	           },
//	       },
//	   },
//	}
//
// errs := validator.Validate(value, nestedValidator)
func (n Field) Validate(value any) error {
	nestedValue, ok := value.(map[string]any)
	if !ok {
		return fmt.Errorf("this value must be a map[string]any")
	}

	var validationErrors []error
	for fieldName, constraint := range n.Constraints {
		fieldValue, exists := nestedValue[fieldName]
		if !exists {
			validationErrors = append(validationErrors, fmt.Errorf("field '%s' is missing", fieldName))
			continue
		}

		if err := constraint.Validate(fieldValue); err != nil {
			validationErrors = append(validationErrors, fmt.Errorf("field '%s': %w", fieldName, err))
		}
	}

	if len(validationErrors) > 0 {
		return fmt.Errorf("validation errors: %v", validationErrors)
	}
	return nil
}
