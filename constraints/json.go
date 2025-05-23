package constraints

import (
	"encoding/json"
	"errors"
)

type IsJSON struct{}

func (c IsJSON) Validate(value any) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("this value should be string type")
	}

	var js any
	if err := json.Unmarshal([]byte(s), &js); err != nil {
		return errors.New("this value should be a valid JSON")
	}

	return nil
}
