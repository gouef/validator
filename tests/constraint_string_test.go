package tests

import (
	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringValidator(t *testing.T) {
	tests := []struct {
		name        string
		value       any
		constraints []validator.Constraint
		expectErr   bool
	}{
		{
			name:        "Valid IsString constraint",
			value:       "abc",
			constraints: []validator.Constraint{constraints.IsString{}},
			expectErr:   false,
		},
		{
			name:        "Invalid IsString constraint",
			value:       35,
			constraints: []validator.Constraint{constraints.IsString{}},
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
