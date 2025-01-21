package tests

import (
	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name        string
		value       any
		constraints []validator.Constraint
		expectErr   bool
	}{
		{
			name:        "Valid Blank constraint",
			value:       "",
			constraints: []validator.Constraint{constraints.Blank{}},
			expectErr:   false,
		},
		{
			name:        "Invalid Blank constraint",
			value:       "some value",
			constraints: []validator.Constraint{constraints.Blank{}},
			expectErr:   true,
		},
		{
			name:        "Valid NotBlank constraint",
			value:       "some value",
			constraints: []validator.Constraint{constraints.NotBlank{}},
			expectErr:   false,
		},
		{
			name:        "Valid (Multiple) NotBlank and IsFalse constraint",
			value:       false,
			constraints: []validator.Constraint{constraints.NotBlank{}, constraints.IsFalse{}},
			expectErr:   false,
		},
		{
			name:        "Invalid NotBlank constraint",
			value:       "",
			constraints: []validator.Constraint{constraints.NotBlank{}},
			expectErr:   true,
		},
		{
			name:        "Valid Nil constraint",
			value:       nil,
			constraints: []validator.Constraint{constraints.Nil{}},
			expectErr:   false,
		},
		{
			name:        "Invalid Nil constraint",
			value:       true,
			constraints: []validator.Constraint{constraints.Nil{}},
			expectErr:   true,
		},
		{
			name:        "Valid NotNil constraint",
			value:       7,
			constraints: []validator.Constraint{constraints.NotNil{}},
			expectErr:   false,
		},
		{
			name:        "Invalid NotNil constraint",
			value:       nil,
			constraints: []validator.Constraint{constraints.NotNil{}},
			expectErr:   true,
		},
		{
			name:        "Valid IsFalse constraint",
			value:       false,
			constraints: []validator.Constraint{constraints.IsFalse{}},
			expectErr:   false,
		},
		{
			name:        "Invalid IsFalse constraint",
			value:       true,
			constraints: []validator.Constraint{constraints.IsFalse{}},
			expectErr:   true,
		},
		{
			name:        "Valid IsTrue constraint",
			value:       true,
			constraints: []validator.Constraint{constraints.IsTrue{}},
			expectErr:   false,
		},
		{
			name:        "Invalid IsTrue constraint",
			value:       false,
			constraints: []validator.Constraint{constraints.IsTrue{}},
			expectErr:   true,
		},
		{
			name:        "Valid IsBool constraint",
			value:       true,
			constraints: []validator.Constraint{constraints.IsBool{}},
			expectErr:   false,
		},
		{
			name:        "Invalid IsBool constraint",
			value:       123,
			constraints: []validator.Constraint{constraints.IsBool{}},
			expectErr:   true,
		},
		{
			name:        "Valid Currency constraint",
			value:       "USD",
			constraints: []validator.Constraint{constraints.Currency{}},
			expectErr:   false,
		},
		{
			name:        "Invalid Currency constraint",
			value:       "INVALID",
			constraints: []validator.Constraint{constraints.Currency{}},
			expectErr:   true,
		},
		{
			name:        "Invalid Currency constraint as type",
			value:       true,
			constraints: []validator.Constraint{constraints.Currency{}},
			expectErr:   true,
		},
		{
			name:        "Valid Email constraint",
			value:       "test@example.com",
			constraints: []validator.Constraint{constraints.Email{}},
			expectErr:   false,
		},
		{
			name:        "Invalid Email constraint",
			value:       "notanemail",
			constraints: []validator.Constraint{constraints.Email{}},
			expectErr:   true,
		},
		{
			name:        "Invalid Email constraint as type",
			value:       true,
			constraints: []validator.Constraint{constraints.Email{}},
			expectErr:   true,
		},
		{
			name:        "Valid Language constraint",
			value:       "en",
			constraints: []validator.Constraint{constraints.Language{}},
			expectErr:   false,
		},
		{
			name:        "Invalid Language constraint",
			value:       "invalidlang",
			constraints: []validator.Constraint{constraints.Language{}},
			expectErr:   true,
		},
		{
			name:        "Invalid Language constraint as type",
			value:       true,
			constraints: []validator.Constraint{constraints.Language{}},
			expectErr:   true,
		},
		{
			name:        "Invalid Language constraint (code)",
			value:       "xx",
			constraints: []validator.Constraint{constraints.Language{}},
			expectErr:   true,
		},
		{
			name:        "Valid LanguageWithRegion constraint",
			value:       "en-US", // Valid language code with region
			constraints: []validator.Constraint{constraints.LanguageWithRegion{}},
			expectErr:   false,
		},
		{
			name:        "Invalid LanguageWithRegion constraint (invalid language code)",
			value:       "invalid-lang", // Invalid language tag
			constraints: []validator.Constraint{constraints.LanguageWithRegion{}},
			expectErr:   true,
		},
		{
			name:        "Invalid LanguageWithRegion constraint (missing region)",
			value:       "en", // Valid language, but missing region
			constraints: []validator.Constraint{constraints.LanguageWithRegion{}},
			expectErr:   true,
		},
		{
			name:        "Invalid LanguageWithRegion constraint (invalid region)",
			value:       "en-XX", // Valid language but invalid region
			constraints: []validator.Constraint{constraints.LanguageWithRegion{}},
			expectErr:   true,
		},
		{
			name:        "Invalid LanguageWithRegion constraint (invalid language and region)",
			value:       "xx-XX", // Valid language but invalid region
			constraints: []validator.Constraint{constraints.LanguageWithRegion{}},
			expectErr:   true,
		},
		{
			name:        "Invalid LanguageWithRegion constraint (invalid type)",
			value:       true, // Valid language but invalid region
			constraints: []validator.Constraint{constraints.LanguageWithRegion{}},
			expectErr:   true,
		},
		{
			name:        "Valid Range constraint",
			value:       5,
			constraints: []validator.Constraint{constraints.RangeConstraint{Min: 1, Max: 10}},
			expectErr:   false,
		},
		{
			name:        "Valid Range constraint (float32)",
			value:       float32(5),
			constraints: []validator.Constraint{constraints.RangeConstraint{Min: 1, Max: 10}},
			expectErr:   false,
		},
		{
			name:        "Valid Range constraint (float64)",
			value:       float64(5),
			constraints: []validator.Constraint{constraints.RangeConstraint{Min: 1, Max: 10}},
			expectErr:   false,
		},
		{
			name:        "Invalid Range constraint",
			value:       15,
			constraints: []validator.Constraint{constraints.RangeConstraint{Min: 1, Max: 10}},
			expectErr:   true,
		},
		{
			name:        "Invalid Range constraint as type",
			value:       "15",
			constraints: []validator.Constraint{constraints.RangeConstraint{Min: 1, Max: 10}},
			expectErr:   true,
		},
		{
			name:        "Valid Required constraint",
			value:       "some value",
			constraints: []validator.Constraint{constraints.Required{}},
			expectErr:   false,
		},
		{
			name:        "Invalid Required constraint",
			value:       "",
			constraints: []validator.Constraint{constraints.Required{}},
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
