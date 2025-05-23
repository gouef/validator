package tests

import (
	"testing"

	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
)

func TestIsJWT(t *testing.T) {
	t.Run("valid JWT", func(t *testing.T) {
		errs := validator.Validate("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.abc.def", constraints.IsJWT{})
		if len(errs) != 0 {
			t.Errorf("unexpected error: %v", errs)
		}
	})

	t.Run("invalid JWT (too few parts)", func(t *testing.T) {
		errs := validator.Validate("only.one", constraints.IsJWT{})
		if len(errs) == 0 {
			t.Errorf("expected error but got none")
		}
	})

	t.Run("invalid JWT (value type)", func(t *testing.T) {
		errs := validator.Validate(127.456, constraints.IsJWT{})
		if len(errs) == 0 {
			t.Errorf("expected error but got none")
		}
	})
}
