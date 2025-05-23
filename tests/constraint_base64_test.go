package tests

import (
	"testing"

	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
)

func TestIsBase64(t *testing.T) {
	t.Run("valid base64", func(t *testing.T) {
		err := validator.Validate("SGVsbG8gd29ybGQ=", constraints.IsBase64{})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("invalid base64", func(t *testing.T) {
		err := validator.Validate("not base64!", constraints.IsBase64{})
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("invalid value", func(t *testing.T) {
		err := validator.Validate(132, constraints.IsBase64{})
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})
}
