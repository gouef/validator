package constraints

import (
	"errors"
	"regexp"
)

type Email struct{}

func (c Email) Validate(value any) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("this value must be a string")
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if matched, _ := regexp.MatchString(emailRegex, str); !matched {
		return errors.New("this value is not a valid email address")
	}

	return nil
}
