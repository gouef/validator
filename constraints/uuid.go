package constraints

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var uuidRegex = regexp.MustCompile(`^[a-f0-9]{8}-[a-f0-9]{4}-[1-5][a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}$`)

type IsUUID struct {
	Version int
}

func (c IsUUID) Validate(value any) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("this value should be string type")
	}

	s = strings.ToLower(s)

	if !uuidRegex.MatchString(s) {
		return errors.New("this value should be a valid UUID")
	}

	if c.Version != 0 {
		if c.Version < 1 || c.Version > 5 || c.Version == 2 {
			return fmt.Errorf("unsupported UUID version: %d", c.Version)
		}
		versionChar := s[14]
		if versionChar != ('0' + byte(c.Version)) {
			return fmt.Errorf("this UUID is not version %d", c.Version)
		}
	}

	return nil
}

type IsUUIDv1 struct{}

func (c IsUUIDv1) Validate(v any) error { return IsUUID{Version: 1}.Validate(v) }

type IsUUIDv2 struct{}

func (c IsUUIDv2) Validate(v any) error { return IsUUID{Version: 2}.Validate(v) }

type IsUUIDv3 struct{}

func (c IsUUIDv3) Validate(v any) error { return IsUUID{Version: 3}.Validate(v) }

type IsUUIDv4 struct{}

func (c IsUUIDv4) Validate(v any) error { return IsUUID{Version: 4}.Validate(v) }

type IsUUIDv5 struct{}

func (c IsUUIDv5) Validate(v any) error { return IsUUID{Version: 5}.Validate(v) }
