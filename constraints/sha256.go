package constraints

import (
	"errors"
	"regexp"
)

type IsSHA256 struct{}

var sha256Regex = regexp.MustCompile(`^[a-fA-F0-9]{64}$`)

func (c IsSHA256) Validate(value any) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("this value should be string type")
	}
	if !sha256Regex.MatchString(s) {
		return errors.New("this value should be a valid SHA256 hash")
	}
	return nil
}
