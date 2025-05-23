package constraints

import (
	"encoding/hex"
	"errors"
)

type IsHex struct{}

func (c IsHex) Validate(value any) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("this value should be string type")
	}
	_, err := hex.DecodeString(s)
	if err != nil {
		return errors.New("this value should be a valid hexadecimal string")
	}
	return nil
}
