package tests

import (
	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegexpValidator(t *testing.T) {
	regExp := `^[a-z]{3}$`
	con := constraints.RegularExpression{
		Regexp: regExp,
	}

	tests := []struct {
		name        string
		value       any
		constraints []validator.Constraint
		expectErr   bool
	}{
		{
			name:        "Valid RegularExpression constraint",
			value:       "abc",
			constraints: []validator.Constraint{con},
			expectErr:   false,
		},
		{
			name:        "Invalid RegularExpression constraint",
			value:       "abcd",
			constraints: []validator.Constraint{con},
			expectErr:   true,
		},
		{
			name:        "Invalid RegularExpression constraint",
			value:       "123",
			constraints: []validator.Constraint{con},
			expectErr:   true,
		},
		{
			name:        "Invalid Struct constraint (not struct)",
			value:       45,
			constraints: []validator.Constraint{con},
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
