package constraints

import (
	"errors"
	"strings"
)

type Contains struct {
	Substring string
}

func (c Contains) Validate(value any) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("this value should be string type")
	}
	if !strings.Contains(s, c.Substring) {
		return errors.New("this value must contain substring: " + c.Substring)
	}
	return nil
}

type HasPrefix struct {
	Prefix string
}

func (c HasPrefix) Validate(value any) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("this value should be string type")
	}
	if !strings.HasPrefix(s, c.Prefix) {
		return errors.New("this value should start with: " + c.Prefix)
	}
	return nil
}

type HasSuffix struct {
	Suffix string
}

func (c HasSuffix) Validate(value any) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("this value should be string type")
	}
	if !strings.HasSuffix(s, c.Suffix) {
		return errors.New("this value should end with: " + c.Suffix)
	}
	return nil
}
