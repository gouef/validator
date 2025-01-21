package tests

import (
	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloatValidator(t *testing.T) {
	tests := []struct {
		name        string
		value       any
		constraints []validator.Constraint
		expectErr   bool
	}{
		{
			name:        "Valid IsFloat constraint",
			value:       45.2,
			constraints: []validator.Constraint{constraints.IsFloat{}},
			expectErr:   false,
		},
		{
			name:        "Valid IsFloat constraint (32)",
			value:       float32(45.2),
			constraints: []validator.Constraint{constraints.IsFloat{}},
			expectErr:   false,
		},
		{
			name:        "Invalid IsFloat constraint",
			value:       "abc",
			constraints: []validator.Constraint{constraints.IsFloat{}},
			expectErr:   true,
		},
		{
			name:        "Valid IsFloat32 constraint",
			value:       float32(45.2),
			constraints: []validator.Constraint{constraints.IsFloat32{}},
			expectErr:   false,
		},
		{
			name:        "Invalid IsFloat32 constraint",
			value:       "abc",
			constraints: []validator.Constraint{constraints.IsFloat32{}},
			expectErr:   true,
		},
		{
			name:        "Valid IsFloat64 constraint",
			value:       45.2,
			constraints: []validator.Constraint{constraints.IsFloat64{}},
			expectErr:   false,
		},
		{
			name:        "Invalid IsFloat64 constraint",
			value:       "abc",
			constraints: []validator.Constraint{constraints.IsFloat64{}},
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
