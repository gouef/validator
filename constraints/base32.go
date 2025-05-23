package constraints

import (
	"encoding/base32"
	"errors"
)

type IsBase32 struct{}

func (c IsBase32) Validate(value any) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("this value should be string type")
	}
	_, err := base32.StdEncoding.DecodeString(s)
	if err != nil {
		return errors.New("this value should be a valid base32 string")
	}
	return nil
}
