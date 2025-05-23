package tests

import (
	"testing"

	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
)

func TestHasSuffix(t *testing.T) {
	t.Run("has suffix", func(t *testing.T) {
		err := validator.Validate("gouef", constraints.HasSuffix{Suffix: "ef"})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("missing suffix", func(t *testing.T) {
		err := validator.Validate("gouef", constraints.HasSuffix{Suffix: "xy"})
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("invalid value", func(t *testing.T) {
		err := validator.Validate(132, constraints.HasSuffix{})
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})
}
