package constraints

import (
	"errors"
	"net/url"
	"strings"
)

type Url struct {
	AllowEmpty     bool
	AllowedSchemes []string
}

// Validate function for validate value
//
// Example:
//
//	v := constraints.Url{
//		AllowEmpty:     false,
//		AllowedSchemes: []string{"http", "https"},
//	}
//
//	value := "https://example.com"
//	errs := validator.Validate(value, v)
func (c Url) Validate(value any) error {
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

	if strValue == "" && c.AllowEmpty {
		return nil
	}

	if strValue == "" {
		return errors.New("this value is required and cannot be empty")
	}

	parsedURL, err := url.ParseRequestURI(strValue)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return errors.New("this value must be a valid URL")
	}

	if len(c.AllowedSchemes) > 0 {
		validScheme := false
		for _, scheme := range c.AllowedSchemes {
			if strings.EqualFold(parsedURL.Scheme, scheme) {
				validScheme = true
				break
			}
		}
		if !validScheme {
			return errors.New("this value must use one of the allowed schemes: " + strings.Join(c.AllowedSchemes, ", "))
		}
	}

	return nil
}
