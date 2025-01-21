package constraints

import (
	"errors"
	"regexp"
)

type RegularExpression struct {
	Regexp string
}

func (c RegularExpression) Validate(value any) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("this value must be a string")
	}

	if matched, _ := regexp.MatchString(c.Regexp, str); !matched {
		return errors.New("this value is not a valid email address")
	}

	return nil
}
