package tests

import (
	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthConstraint(t *testing.T) {
	tests := []struct {
		name          string
		constraint    constraints.Length
		value         any
		expectErr     bool
		expectedError string
	}{
		{"Valid string length", constraints.Length{Min: 3, Max: 10}, "hello", false, ""},
		{"String too short", constraints.Length{Min: 3}, "hi", true, "this value must have a length of at least 3"},
		{"String too long", constraints.Length{Max: 5}, "hello world", true, "this value must have a length of at most 5"},
		{"Valid array length", constraints.Length{Min: 2, Max: 4}, []int{1, 2, 3}, false, ""},
		{"Array too short", constraints.Length{Min: 4}, []int{1, 2, 3}, true, "this value must have a length of at least 4"},
		{"Map too long", constraints.Length{Max: 2}, map[string]int{"a": 1, "b": 2, "c": 3}, true, "this value must have a length of at most 2"},
		{"Valid empty slice", constraints.Length{Min: 0, Max: 0}, []int{}, false, ""},
		{"Invalid type", constraints.Length{Min: 3, Max: 5}, 12345, true, "this value must be a string, array, slice, or map"},
		{"Nil value", constraints.Length{Min: 1}, nil, true, "this value is required and cannot be nil"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			errs := validator.Validate(test.value, test.constraint)

			if test.expectErr {
				assert.NotEmpty(t, errs, "Expected error but got none")
				if len(errs) > 0 {
					assert.EqualError(t, errs[0], test.expectedError)
				}
			} else {
				assert.Empty(t, errs, "Expected no error but got some")
			}
		})
	}
}
