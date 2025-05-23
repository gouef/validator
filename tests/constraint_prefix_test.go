package tests

import (
	"testing"

	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
)

func TestHasPrefix(t *testing.T) {
	t.Run("has prefix", func(t *testing.T) {
		err := validator.Validate("gouef", constraints.HasPrefix{Prefix: "go"})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("missing prefix", func(t *testing.T) {
		err := validator.Validate("gouef", constraints.HasPrefix{Prefix: "mo"})
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("invalid value", func(t *testing.T) {
		err := validator.Validate(132, constraints.HasPrefix{})
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})
}
