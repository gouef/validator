package constraints

import (
	"encoding/base64"
	"errors"
)

type IsBase64 struct{}

func (c IsBase64) Validate(value any) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("this value should be string type")
	}
	_, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return errors.New("this value should be a valid base64 string")
	}
	return nil
}
