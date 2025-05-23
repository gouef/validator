package tests

import (
	"errors"
	"testing"

	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
)

func TestCustomConstraint(t *testing.T) {
	t.Run("valid custom rule", func(t *testing.T) {
		con := constraints.Custom{
			Fn: func(v any) error {
				s, ok := v.(string)
				if !ok {
					return errors.New("not string")
				}
				if len(s) < 3 {
					return errors.New("too short")
				}
				return nil
			},
		}

		errs := validator.Validate("hello", con)
		if len(errs) != 0 {
			t.Errorf("unexpected error: %v", errs)
		}
	})

	t.Run("invalid custom rule", func(t *testing.T) {
		con := constraints.Custom{
			Fn: func(v any) error {
				return errors.New("always fail")
			},
		}

		errs := validator.Validate("anything", con)
		if len(errs) == 0 {
			t.Errorf("expected error but got none")
		}
	})

	t.Run("missing function", func(t *testing.T) {
		con := constraints.Custom{} // Fn == nil
		errs := validator.Validate("test", con)
		if len(errs) == 0 || errs[0].Error() != "custom constraint: missing function" {
			t.Errorf("expected error for nil function")
		}
	})
}
