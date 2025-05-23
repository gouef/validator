package tests

import (
	"testing"

	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
)

func TestIsHex(t *testing.T) {
	t.Run("valid hex", func(t *testing.T) {
		err := validator.Validate("deadbeef", constraints.IsHex{})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("invalid hex", func(t *testing.T) {
		err := validator.Validate("not-hex", constraints.IsHex{})
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("invalid value", func(t *testing.T) {
		err := validator.Validate(132, constraints.IsHex{})
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})
}
