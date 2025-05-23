package tests

import (
	"testing"

	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
)

func TestIsHSLColor(t *testing.T) {
	t.Run("valid hsl", func(t *testing.T) {
		err := validator.Validate("hsl(120, 100%, 50%)", constraints.IsHSLColor{})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("valid hsla", func(t *testing.T) {
		err := validator.Validate("hsla(120, 100%, 50%, 0.5)", constraints.IsHSLColor{})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("invalid hsl", func(t *testing.T) {
		err := validator.Validate("hsl(120, 100, 50)", constraints.IsHSLColor{})
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("invalid format", func(t *testing.T) {
		err := validator.Validate("not-a-color", constraints.IsHSLColor{})
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("invalid hex (bad value)", func(t *testing.T) {
		err := validator.Validate(135, constraints.IsHSLColor{})
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})
}
