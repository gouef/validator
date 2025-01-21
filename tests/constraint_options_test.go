package tests

import (
	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionsValidator(t *testing.T) {
	validatorConstrain := constraints.Options{
		Options: []any{"option1", "option2", "option3", 42},
	}

	arrayValidatorConstrain := constraints.ArrayOptions{
		Options: []any{"Red", "Green", "Blue", 100},
	}

	tests := []struct {
		name        string
		value       any
		constraints []validator.Constraint
		expectErr   bool
	}{
		{
			name:        "Valid Options constraint",
			value:       "option1",
			constraints: []validator.Constraint{validatorConstrain},
			expectErr:   false,
		},
		{
			name:        "Valid Options constraint (int)",
			value:       42,
			constraints: []validator.Constraint{validatorConstrain},
			expectErr:   false,
		},
		{
			name:        "Invalid Options constraint (invalid)",
			value:       "invalid",
			constraints: []validator.Constraint{validatorConstrain},
			expectErr:   true,
		},
		{
			name:        "Invalid Options constraint (empty)",
			value:       "",
			constraints: []validator.Constraint{validatorConstrain},
			expectErr:   true,
		},
		{
			name:        "Invalid Options constraint (nil)",
			value:       nil,
			constraints: []validator.Constraint{validatorConstrain},
			expectErr:   true,
		},
		{
			name:        "Valid ArrayOptions constraint",
			value:       []any{"Red", "Green", 100},
			constraints: []validator.Constraint{arrayValidatorConstrain},
			expectErr:   false,
		},
		{
			name:        "Valid ArrayOptions constraint (empty)",
			value:       []any{},
			constraints: []validator.Constraint{arrayValidatorConstrain},
			expectErr:   false,
		},
		{
			name:        "Invalid ArrayOptions constraint (one element)",
			value:       []any{"Red", "Yellow"},
			constraints: []validator.Constraint{arrayValidatorConstrain},
			expectErr:   true,
		},
		{
			name:        "Invalid ArrayOptions constraint (multiple elements)",
			value:       []any{"Yellow", "Purple"},
			constraints: []validator.Constraint{arrayValidatorConstrain},
			expectErr:   true,
		},
		{
			name:        "Invalid ArrayOptions constraint (not array)",
			value:       nil,
			constraints: []validator.Constraint{arrayValidatorConstrain},
			expectErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := validator.Validate(tt.value, tt.constraints...)
			if tt.expectErr {
				assert.NotEmpty(t, errs, "Expected error but got none")
			} else {
				assert.Empty(t, errs, "Expected no error but got some")
			}
		})
	}
}
