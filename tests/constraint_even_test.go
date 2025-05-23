package tests

import (
	"testing"

	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
)

func TestIsEven(t *testing.T) {
	t.Run("even number", func(t *testing.T) {
		errs := validator.Validate(4, constraints.IsEven{})
		if len(errs) != 0 {
			t.Errorf("unexpected error: %v", errs)
		}
	})

	t.Run("odd number", func(t *testing.T) {
		errs := validator.Validate(3, constraints.IsEven{})
		if len(errs) == 0 {
			t.Errorf("expected error but got none")
		}
	})

	t.Run("invalid type", func(t *testing.T) {
		errs := validator.Validate("test", constraints.IsEven{})
		if len(errs) == 0 {
			t.Errorf("expected division by zero error")
		}
	})
}
