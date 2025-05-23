package tests

import (
	"testing"

	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
)

func TestIsBase32(t *testing.T) {
	t.Run("valid base32", func(t *testing.T) {
		errs := validator.Validate("MFRGGZDFMZTWQ2LK", constraints.IsBase32{})
		if len(errs) != 0 {
			t.Errorf("unexpected error: %v", errs)
		}
	})

	t.Run("invalid base32", func(t *testing.T) {
		errs := validator.Validate("not!base32", constraints.IsBase32{})
		if len(errs) == 0 {
			t.Errorf("expected error but got none")
		}
	})

	t.Run("invalid base32 value type", func(t *testing.T) {
		errs := validator.Validate(32, constraints.IsBase32{})
		if len(errs) == 0 {
			t.Errorf("expected error but got none")
		}
	})
}
