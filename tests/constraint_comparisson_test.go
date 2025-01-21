package tests

import (
	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompareValidate(t *testing.T) {
	tests := []struct {
		name        string
		value       any
		constraints []validator.Constraint
		expectErr   bool
	}{
		{
			name:        "Valid EqualTo constraint",
			value:       "test",
			constraints: []validator.Constraint{constraints.EqualTo{Value: "test"}},
			expectErr:   false,
		},
		{
			name:        "Valid EqualTo constraint (int)",
			value:       5,
			constraints: []validator.Constraint{constraints.EqualTo{Value: 5}},
			expectErr:   false,
		},
		{
			name:        "Invalid EqualTo constraint",
			value:       5,
			constraints: []validator.Constraint{constraints.EqualTo{Value: 6}},
			expectErr:   true,
		},
		{
			name:        "Valid GreaterThen constraint)",
			value:       6,
			constraints: []validator.Constraint{constraints.GreaterThen{Value: 5}},
			expectErr:   false,
		},
		{
			name:        "Invalid GreaterThen constraint",
			value:       5,
			constraints: []validator.Constraint{constraints.GreaterThen{Value: 6}},
			expectErr:   true,
		},
		{
			name:        "Invalid GreaterThen constraint (type)",
			value:       "5",
			constraints: []validator.Constraint{constraints.GreaterThen{Value: 6}},
			expectErr:   true,
		},
		{
			name:        "Valid GreaterOrEqual constraint",
			value:       6,
			constraints: []validator.Constraint{constraints.GreaterOrEqual{Value: 5}},
			expectErr:   false,
		},
		{
			name:        "Valid GreaterThen constraint (equal)",
			value:       5,
			constraints: []validator.Constraint{constraints.GreaterOrEqual{Value: 5}},
			expectErr:   false,
		},
		{
			name:        "Invalid GreaterOrEqual constraint",
			value:       5,
			constraints: []validator.Constraint{constraints.GreaterThen{Value: 6}},
			expectErr:   true,
		},
		{
			name:        "Invalid GreaterOrEqual constraint (type)",
			value:       "5",
			constraints: []validator.Constraint{constraints.GreaterThen{Value: 6}},
			expectErr:   true,
		},
		{
			name:        "Valid LessOrEqual constraint",
			value:       4,
			constraints: []validator.Constraint{constraints.LessOrEqual{Value: 5}},
			expectErr:   false,
		},
		{
			name:        "Valid LessOrEqual constraint (equal)",
			value:       5,
			constraints: []validator.Constraint{constraints.LessOrEqual{Value: 5}},
			expectErr:   false,
		},
		{
			name:        "Invalid LessOrEqual constraint",
			value:       7,
			constraints: []validator.Constraint{constraints.LessOrEqual{Value: 6}},
			expectErr:   true,
		},
		{
			name:        "Invalid LessOrEqual constraint (type)",
			value:       "7",
			constraints: []validator.Constraint{constraints.LessOrEqual{Value: 6}},
			expectErr:   true,
		},
		{
			name:        "Valid LessThen constraint",
			value:       4,
			constraints: []validator.Constraint{constraints.LessThen{Value: 5}},
			expectErr:   false,
		},
		{
			name:        "Invalid LessThen constraint",
			value:       7,
			constraints: []validator.Constraint{constraints.LessThen{Value: 6}},
			expectErr:   true,
		},
		{
			name:        "Invalid LessThen constraint (type)",
			value:       "7",
			constraints: []validator.Constraint{constraints.LessThen{Value: 6}},
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
