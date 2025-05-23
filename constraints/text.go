package constraints

import (
	"errors"
	"strings"
	"unicode"
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

type IsLowercase struct{}

func (c IsLowercase) Validate(value any) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("this value should be string type")
	}
	if s != strings.ToLower(s) {
		return errors.New("this value should be all lowercase")
	}
	return nil
}

type IsUppercase struct{}

func (c IsUppercase) Validate(value any) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("this value should be string type")
	}
	if s != strings.ToUpper(s) {
		return errors.New("this value should be all uppercase")
	}
	return nil
}

type IsTitleCase struct{}

func (c IsTitleCase) Validate(value any) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("this value should be string type")
	}
	if len(s) == 0 || !unicode.IsUpper([]rune(s)[0]) {
		return errors.New("this value should start with an uppercase letter")
	}
	return nil
}

type IsCamelCase struct{}

func (c IsCamelCase) Validate(value any) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("this value should be string type")
	}
	r := []rune(s)
	if len(r) == 0 || !unicode.IsLower(r[0]) {
		return errors.New("this value should be in camelCase")
	}
	for _, ch := range r {
		if ch == ' ' || ch == '_' {
			return errors.New("this value should not contain spaces or underscores")
		}
	}
	return nil
}

type IsPascalCase struct{}

func (c IsPascalCase) Validate(value any) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("this value should be string type")
	}
	r := []rune(s)
	if len(r) == 0 || !unicode.IsUpper(r[0]) {
		return errors.New("this value should be in PascalCase")
	}
	for _, ch := range r {
		if ch == ' ' || ch == '_' {
			return errors.New("this value should not contain spaces or underscores")
		}
	}
	return nil
}
