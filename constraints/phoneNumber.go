package constraints

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type PhoneNumber struct {
	Format     string // Format string, e.g., "(00420) 000 000 000"
	AllowEmpty bool   // Allow empty values
}

// Helper function to convert format to regex
func formatToRegex(format string) string {
	var regex strings.Builder
	for _, ch := range format {
		switch ch {
		case '0':
			regex.WriteString(`\d`)
		case ' ':
			regex.WriteString(`\s`)
		default:
			regex.WriteString(regexp.QuoteMeta(string(ch)))
		}
	}
	return "^" + regex.String() + "$"
}

// Validate function for validate value
//
// Example:
//
//	con := constraints.PhoneNumber{Format: "(00420) 000 000 000"}
//
//	value := "(00420) 123 456 789"
//	errs := validator.Validate(value, v)
func (c PhoneNumber) Validate(value any) error {
	if value == nil {
		if c.AllowEmpty {
			return nil
		}
		return errors.New("this value is required and cannot be nil")
	}

	strValue, ok := value.(string)
	if !ok {
		return errors.New("this value must be a string")
	}

	if strValue == "" {
		if c.AllowEmpty {
			return nil
		}
		return errors.New("this value is required and cannot be empty")
	}

	regex := regexp.MustCompile(formatToRegex(c.Format))

	if !regex.MatchString(strValue) {
		return fmt.Errorf("this value must match the phone number format: %s", c.Format)
	}

	return nil
}
