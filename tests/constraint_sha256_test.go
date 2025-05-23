package tests

import (
	"testing"

	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
)

func TestIsSHA256(t *testing.T) {
	t.Run("valid hash", func(t *testing.T) {
		errs := validator.Validate("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", constraints.IsSHA256{})
		if len(errs) != 0 {
			t.Errorf("unexpected error: %v", errs)
		}
	})

	t.Run("invalid length", func(t *testing.T) {
		errs := validator.Validate("abc123", constraints.IsSHA256{})
		if len(errs) == 0 {
			t.Errorf("expected error but got none")
		}
	})

	t.Run("invalid value type", func(t *testing.T) {
		errs := validator.Validate(123, constraints.IsSHA256{})
		if len(errs) == 0 {
			t.Errorf("expected error but got none")
		}
	})
}
