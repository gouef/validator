package tests

import (
	"fmt"
	"testing"

	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
	"github.com/stretchr/testify/assert"
)

func TestURLConstraint(t *testing.T) {
	v := constraints.Url{
		AllowEmpty:     false,
		AllowedSchemes: []string{"http", "https"},
	}

	tests := []struct {
		name      string
		value     any
		expectErr bool
	}{
		{"Valid URL", "https://example.com", false},
		{"Valid URL with path", "https://example.com/path?query=1", false},
		{"Invalid URL", "htp:/example", true},
		{"Invalid URL scheme", "ftp://example.com", true},
		{"Empty value (not allowed)", "", true},
		{"Nil value (not allowed)", nil, true},
		{"Empty value (allowed)", "", false},
		{"Nil value (allowed)", nil, false},
		{"Invalid type", 12345, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.name == "Empty value (allowed)" || test.name == "Nil value (allowed)" {
				v.AllowEmpty = true
			}
			errs := validator.Validate(test.value, v)

			if test.expectErr {
				assert.NotEmpty(t, errs, fmt.Sprintf("[Test: %s] Expected error but got none", test.name))
			} else {
				assert.Empty(t, errs, fmt.Sprintf("[Test: %s] Expected no error but got some", test.name))
			}
		})
	}
}
