package constraints

import (
	"errors"
	"strings"
)

type IsJWT struct{}

func (c IsJWT) Validate(value any) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("this value should be string type")
	}

	parts := strings.Split(s, ".")
	if len(parts) != 3 {
		return errors.New("this value should be a valid JWT (three segments)")
	}

	return nil
}
