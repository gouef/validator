package tests

import (
	"testing"

	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
)

func TestIsJSON(t *testing.T) {
	t.Run("valid JSON object", func(t *testing.T) {
		errs := validator.Validate(`{"name":"John","age":30}`, constraints.IsJSON{})
		if len(errs) != 0 {
			t.Errorf("unexpected error: %v", errs)
		}
	})

	t.Run("valid JSON array", func(t *testing.T) {
		errs := validator.Validate(`[1, 2, 3]`, constraints.IsJSON{})
		if len(errs) != 0 {
			t.Errorf("unexpected error: %v", errs)
		}
	})

	t.Run("valid JSON string", func(t *testing.T) {
		errs := validator.Validate(`"hello"`, constraints.IsJSON{})
		if len(errs) != 0 {
			t.Errorf("unexpected error: %v", errs)
		}
	})

	t.Run("invalid JSON", func(t *testing.T) {
		errs := validator.Validate(`{name: John}`, constraints.IsJSON{})
		if len(errs) == 0 {
			t.Errorf("expected error but got none")
		}
	})

	t.Run("non-string value", func(t *testing.T) {
		errs := validator.Validate(123, constraints.IsJSON{})
		if len(errs) == 0 {
			t.Errorf("expected error but got none")
		}
	})
}
