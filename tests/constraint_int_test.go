package tests

import (
	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntValidator(t *testing.T) {
	tests := []struct {
		name        string
		value       any
		constraints []validator.Constraint
		expectErr   bool
	}{
		{
			name:        "Valid IsInt constraint",
			value:       45,
			constraints: []validator.Constraint{constraints.IsInt{}},
			expectErr:   false,
		},
		{
			name:        "Invalid IsInt constraint",
			value:       "abc",
			constraints: []validator.Constraint{constraints.IsInt{}},
			expectErr:   true,
		},
		{
			name:        "Valid IsInt8 constraint",
			value:       int8(45),
			constraints: []validator.Constraint{constraints.IsInt8{}},
			expectErr:   false,
		},
		{
			name:        "Invalid IsInt8 constraint",
			value:       "abc",
			constraints: []validator.Constraint{constraints.IsInt8{}},
			expectErr:   true,
		},
		{
			name:        "Valid IsInt16 constraint",
			value:       int16(45),
			constraints: []validator.Constraint{constraints.IsInt16{}},
			expectErr:   false,
		},
		{
			name:        "Invalid IsInt16 constraint",
			value:       "abc",
			constraints: []validator.Constraint{constraints.IsInt16{}},
			expectErr:   true,
		},
		{
			name:        "Valid IsInt32 constraint",
			value:       int32(45),
			constraints: []validator.Constraint{constraints.IsInt32{}},
			expectErr:   false,
		},
		{
			name:        "Invalid IsInt32 constraint",
			value:       "abc",
			constraints: []validator.Constraint{constraints.IsInt32{}},
			expectErr:   true,
		},
		{
			name:        "Valid IsInt64 constraint",
			value:       int64(45),
			constraints: []validator.Constraint{constraints.IsInt64{}},
			expectErr:   false,
		},
		{
			name:        "Invalid IsInt64 constraint",
			value:       "abc",
			constraints: []validator.Constraint{constraints.IsInt64{}},
			expectErr:   true,
		},
		{
			name:        "Valid IsUint constraint",
			value:       uint(45),
			constraints: []validator.Constraint{constraints.IsUint{}},
			expectErr:   false,
		},
		{
			name:        "Invalid IsUint constraint",
			value:       "abc",
			constraints: []validator.Constraint{constraints.IsUint{}},
			expectErr:   true,
		},
		{
			name:        "Valid IsUint8 constraint",
			value:       uint8(45),
			constraints: []validator.Constraint{constraints.IsUint8{}},
			expectErr:   false,
		},
		{
			name:        "Invalid IsUint8 constraint",
			value:       "abc",
			constraints: []validator.Constraint{constraints.IsUint8{}},
			expectErr:   true,
		},
		{
			name:        "Valid IsUint16 constraint",
			value:       uint16(45),
			constraints: []validator.Constraint{constraints.IsUint16{}},
			expectErr:   false,
		},
		{
			name:        "Invalid IsUint16 constraint",
			value:       "abc",
			constraints: []validator.Constraint{constraints.IsUint16{}},
			expectErr:   true,
		},
		{
			name:        "Valid IsUint32 constraint",
			value:       uint32(45),
			constraints: []validator.Constraint{constraints.IsUint32{}},
			expectErr:   false,
		},
		{
			name:        "Invalid IsUint32 constraint",
			value:       "abc",
			constraints: []validator.Constraint{constraints.IsUint32{}},
			expectErr:   true,
		},
		{
			name:        "Valid IsUint64 constraint",
			value:       uint64(45),
			constraints: []validator.Constraint{constraints.IsUint64{}},
			expectErr:   false,
		},
		{
			name:        "Invalid IsUint64 constraint",
			value:       "abc",
			constraints: []validator.Constraint{constraints.IsUint64{}},
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
