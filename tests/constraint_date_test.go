package tests

import (
	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	dateFormat := "2006-01-02"
	minDate, _ := time.Parse(dateFormat, "2023-01-01")
	maxDate, _ := time.Parse(dateFormat, "2023-12-31")

	con := constraints.Date{
		Min:    minDate,
		Max:    maxDate,
		Format: dateFormat,
	}

	tests := []struct {
		name      string
		value     any
		expectErr bool
	}{
		{"Valid time.Time date", time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC), false},
		{"Valid string date", "2023-05-01", false},
		{"Date before range", "2022-12-31", true},
		{"Date after range", "2024-01-01", true},
		{"Invalid date format", "01-05-2023", true},
		{"Nil value (not allowed)", nil, true},
		{"Nil value (allowed)", nil, false},
		{"Invalid type", 12345, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.name == "Nil value (allowed)" {
				con.AllowNil = true
			}
			errs := validator.Validate(test.value, con)

			if test.expectErr {
				assert.NotEmpty(t, errs, "Expected error but got none")
			} else {
				assert.Empty(t, errs, "Expected no error but got some")
			}
		})
	}

	t.Run("Invalid format", func(t *testing.T) {
		dateFormat = ""
		minDate, _ = time.Parse(dateFormat, "2023-01-01")
		maxDate, _ = time.Parse(dateFormat, "2023-12-31")

		con = constraints.Date{
			Min:      minDate,
			Max:      maxDate,
			Format:   dateFormat,
			AllowNil: true,
		}

		errs := validator.Validate("2023-05-01", con)
		assert.NotEmpty(t, errs, "Expected error but got none")
	})
}
