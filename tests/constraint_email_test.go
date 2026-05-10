package tests

import (
	"testing"

	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {

	con := constraints.Email{}

	tests := []struct {
		name         string
		value        any
		expReturnVal bool
	}{
		{"Valid", "some@email.com", true},
		{"Invalid type", 12345, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			errs := validator.ValidateOk(test.value, con)

			assert.Equal(t, test.expReturnVal, errs, "Expected validate be ok")
		})
	}

}
