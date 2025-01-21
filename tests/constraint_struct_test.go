package tests

import (
	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MyStruct struct {
	Field1 string
	Field2 int
}

type InvalidStruct struct {
	Field1 string
}

func TestStructValidator(t *testing.T) {
	validStruct := MyStruct{
		Field1: "Valid",
		Field2: 123,
	}

	invalidStruct := InvalidStruct{
		Field1: "Invalid",
	}

	invalidStruct2 := struct {
		Field2 int
	}{
		Field2: 456,
	}

	con := constraints.Struct{
		Struct: MyStruct{},
	}

	tests := []struct {
		name        string
		value       any
		constraints []validator.Constraint
		expectErr   bool
	}{
		{
			name:        "Valid Struct constraint",
			value:       validStruct,
			constraints: []validator.Constraint{con},
			expectErr:   false,
		},
		{
			name:        "Invalid Struct constraint",
			value:       invalidStruct,
			constraints: []validator.Constraint{con},
			expectErr:   true,
		},
		{
			name:        "Invalid Struct 2 constraint",
			value:       invalidStruct2,
			constraints: []validator.Constraint{con},
			expectErr:   true,
		},
		{
			name:        "Invalid Struct constraint (not struct)",
			value:       "string",
			constraints: []validator.Constraint{con},
			expectErr:   true,
		},
		{
			name:        "Valid IsStruct constraint",
			value:       validStruct,
			constraints: []validator.Constraint{constraints.IsStruct{}},
			expectErr:   false,
		},
		{
			name:        "Invalid IsStruct constraint",
			value:       65,
			constraints: []validator.Constraint{constraints.IsStruct{}},
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
