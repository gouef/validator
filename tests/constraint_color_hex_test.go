package tests

import (
	"testing"

	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
)

func TestIsHexColor(t *testing.T) {
	t.Run("valid short hex", func(t *testing.T) {
		err := validator.Validate("#fff", constraints.IsHexColor{})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("valid full hex", func(t *testing.T) {
		err := validator.Validate("#ff00cc", constraints.IsHexColor{})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("invalid hex (no hash)", func(t *testing.T) {
		err := validator.Validate("ff00cc", constraints.IsHexColor{})
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("invalid hex (bad length)", func(t *testing.T) {
		err := validator.Validate("#ff0c", constraints.IsHexColor{})
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("invalid hex (bad value)", func(t *testing.T) {
		err := validator.Validate(135, constraints.IsHexColor{})
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})
}
