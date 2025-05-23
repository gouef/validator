package tests

import (
	"testing"

	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
)

func TestIsMultipleOf(t *testing.T) {
	t.Run("multiple of 5", func(t *testing.T) {
		errs := validator.Validate(25, constraints.IsMultipleOf{Divisor: 5})
		if len(errs) != 0 {
			t.Errorf("unexpected error: %v", errs)
		}
	})

	t.Run("not a multiple", func(t *testing.T) {
		errs := validator.Validate(7, constraints.IsMultipleOf{Divisor: 3})
		if len(errs) == 0 {
			t.Errorf("expected error but got none")
		}
	})

	t.Run("zero divisor", func(t *testing.T) {
		errs := validator.Validate(7, constraints.IsMultipleOf{Divisor: 0})
		if len(errs) == 0 || errs[0].Error() != "divisor cannot be zero" {
			t.Errorf("expected division by zero error")
		}
	})

	t.Run("invalid type", func(t *testing.T) {
		errs := validator.Validate("test", constraints.IsMultipleOf{Divisor: 0})
		if len(errs) == 0 {
			t.Errorf("expected division by zero error")
		}
	})
}
