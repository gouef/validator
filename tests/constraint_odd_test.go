package tests

import (
	"testing"

	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
)

func TestIsOdd(t *testing.T) {
	t.Run("odd number", func(t *testing.T) {
		errs := validator.Validate(5, constraints.IsOdd{})
		if len(errs) != 0 {
			t.Errorf("unexpected error: %v", errs)
		}
	})

	t.Run("even number", func(t *testing.T) {
		errs := validator.Validate(2, constraints.IsOdd{})
		if len(errs) == 0 {
			t.Errorf("expected error but got none")
		}
	})

	t.Run("invalid type", func(t *testing.T) {
		errs := validator.Validate("test", constraints.IsOdd{})
		if len(errs) == 0 {
			t.Errorf("expected division by zero error")
		}
	})
}
