package tests

import (
	"fmt"
	"github.com/gouef/validator/constraints"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountry_Validate(t *testing.T) {
	tests := []struct {
		value    any
		expected error
	}{
		{
			value:    "CZ",
			expected: nil,
		},
		{
			value:    "CZE",
			expected: nil,
		},
		{
			value:    "Czech Republic",
			expected: fmt.Errorf("invalid country: Czech Republic"),
		},
		{
			value:    "Czechia",
			expected: nil,
		},
		{
			value:    "ZZ",
			expected: fmt.Errorf("invalid country: ZZ"),
		},
		{
			value:    "SomeInvalidCountry",
			expected: fmt.Errorf("invalid country: SomeInvalidCountry"),
		},
		{
			value:    123,
			expected: fmt.Errorf("this value must be a string"),
		},
	}

	con := constraints.Country{}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("testing %v", tt.value), func(t *testing.T) {
			err := con.Validate(tt.value)
			if tt.expected == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.expected.Error())
			}
		})
	}
}
