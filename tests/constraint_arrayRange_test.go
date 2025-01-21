package tests

import (
	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayRangeValidator(t *testing.T) {
	validatorConstrain := constraints.ArrayRangeConstraint{
		Min: 10,
		Max: 20,
	}

	tests := []struct {
		name        string
		value       any
		constraints []validator.Constraint
		expectErr   bool
	}{
		{
			name:        "Valid ArrayRangeConstraint constraint",
			value:       []any{10, 15, 20},
			constraints: []validator.Constraint{validatorConstrain},
			expectErr:   false,
		},
		{
			name:        "Valid ArrayRangeConstraint constraint (float)",
			value:       []any{float32(10), float64(15), 20},
			constraints: []validator.Constraint{validatorConstrain},
			expectErr:   false,
		},
		{
			name:        "Valid ArrayRangeConstraint constraint (empty)",
			value:       []any{},
			constraints: []validator.Constraint{validatorConstrain},
			expectErr:   false,
		},
		{
			name:        "Invalid ArrayRangeConstraint constraint (out of range)",
			value:       []any{5, 15, 25},
			constraints: []validator.Constraint{validatorConstrain},
			expectErr:   true,
		},
		{
			name:        "Invalid ArrayRangeConstraint constraint (mixed and invalid)",
			value:       []any{10, "invalid", 15},
			constraints: []validator.Constraint{validatorConstrain},
			expectErr:   true,
		},
		{
			name:        "Invalid ArrayRangeConstraint constraint (not array)",
			value:       "abc",
			constraints: []validator.Constraint{validatorConstrain},
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
