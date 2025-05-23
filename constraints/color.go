package constraints

import (
	"errors"
	"regexp"
)

type IsHexColor struct{}

var hexColorRegex = regexp.MustCompile(`^#(?:[0-9a-fA-F]{3}|[0-9a-fA-F]{6})$`)

func (c IsHexColor) Validate(value any) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("this value should be string type")
	}
	if !hexColorRegex.MatchString(s) {
		return errors.New("this value should be a valid hex color (e.g. #fff or #ffffff)")
	}
	return nil
}

type IsRGBColor struct{}

var rgbColorRegex = regexp.MustCompile(`(?i)^rgba?\(\s*\d{1,3}\s*,\s*\d{1,3}\s*,\s*\d{1,3}(?:\s*,\s*(0|0?\.\d+|1(\.0)?))?\s*\)$`)

func (c IsRGBColor) Validate(value any) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("this value should be string type")
	}
	if !rgbColorRegex.MatchString(s) {
		return errors.New("this value should be a valid rgb or rgba color")
	}
	return nil
}

type IsHSLColor struct{}

var hslColorRegex = regexp.MustCompile(`(?i)^hsla?\(\s*\d{1,3}\s*,\s*\d{1,3}%\s*,\s*\d{1,3}%(\s*,\s*(0|0?\.\d+|1(\.0)?))?\s*\)$`)

func (c IsHSLColor) Validate(value any) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("this value should be string type")
	}
	if !hslColorRegex.MatchString(s) {
		return errors.New("this value should be a valid hsl or hsla color")
	}
	return nil
}
