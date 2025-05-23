package tests

import (
	"testing"

	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
)

func TestIsRGBColor(t *testing.T) {
	t.Run("valid rgb", func(t *testing.T) {
		err := validator.Validate("rgb(255, 0, 127)", constraints.IsRGBColor{})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("valid rgba", func(t *testing.T) {
		err := validator.Validate("rgba(255, 0, 127, 0.5)", constraints.IsRGBColor{})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("invalid rgb (missing parens)", func(t *testing.T) {
		err := validator.Validate("rgb 255,0,0", constraints.IsRGBColor{})
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("invalid rgba (alpha out of range)", func(t *testing.T) {
		err := validator.Validate("rgba(255, 0, 0, 2)", constraints.IsRGBColor{})
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("invalid hex (bad value)", func(t *testing.T) {
		err := validator.Validate(135, constraints.IsRGBColor{})
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})
}
