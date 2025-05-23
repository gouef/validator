package tests

import (
	"testing"

	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
)

func TestContains(t *testing.T) {
	t.Run("contains substring", func(t *testing.T) {
		err := validator.Validate("hello world", constraints.Contains{Substring: "world"})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("does not contain substring", func(t *testing.T) {
		err := validator.Validate("hello world", constraints.Contains{Substring: "mars"})
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("invalid value", func(t *testing.T) {
		err := validator.Validate(132, constraints.Contains{Substring: "mars"})
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})
}
