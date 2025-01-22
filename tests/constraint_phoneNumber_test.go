package tests

import (
	"testing"

	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
	"github.com/stretchr/testify/assert"
)

func TestPhoneNumberConstraint(t *testing.T) {
	tests := []struct {
		name          string
		constraint    constraints.PhoneNumber
		value         any
		expectErr     bool
		expectedError string
	}{
		{"Valid phone number", constraints.PhoneNumber{Format: "(00420) 000 000 000"}, "(00420) 123 456 789", false, ""},
		{"Invalid phone number", constraints.PhoneNumber{Format: "(00420) 000 000 000"}, "123456789", true, "this value must match the phone number format"},
		{"Empty value allowed", constraints.PhoneNumber{Format: "(00420) 000 000 000", AllowEmpty: true}, "", false, ""},
		{"Empty value not allowed", constraints.PhoneNumber{Format: "(00420) 000 000 000", AllowEmpty: false}, "", true, "this value is required and cannot be empty"},
		{"Nil value allowed", constraints.PhoneNumber{Format: "(00420) 000 000 000", AllowEmpty: true}, nil, false, ""},
		{"Nil value not allowed", constraints.PhoneNumber{Format: "(00420) 000 000 000", AllowEmpty: false}, nil, true, "this value is required and cannot be nil"},
		{"not string value not allowed", constraints.PhoneNumber{Format: "(00420) 000 000 000", AllowEmpty: false}, 45, true, "this value must be"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			errs := validator.Validate(test.value, test.constraint)

			if test.expectErr {
				assert.NotEmpty(t, errs, "Expected error but got none")
				if len(errs) > 0 {
					assert.Contains(t, errs[0].Error(), test.expectedError)
				}
			} else {
				assert.Empty(t, errs, "Expected no error but got some")
			}
		})
	}
}
