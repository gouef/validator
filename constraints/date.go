package constraints

import (
	"errors"
	"fmt"
	"time"
)

type Date struct {
	Min      time.Time
	Max      time.Time
	Format   string
	AllowNil bool
}

func (c Date) Validate(value any) error {
	if value == nil {
		if c.AllowNil {
			return nil
		}
		return errors.New("this value is required and cannot be nil")
	}

	var date time.Time
	switch v := value.(type) {
	case time.Time:
		date = v
	case string:
		if c.Format == "" {
			return errors.New("date format is required for string values")
		}
		parsed, err := time.Parse(c.Format, v)
		if err != nil {
			return fmt.Errorf("invalid date format: expected %s", c.Format)
		}
		date = parsed
	default:
		return errors.New("this value must be a valid date (time.Time or string)")
	}

	if !date.IsZero() && (date.Before(c.Min) || date.After(c.Max)) {
		return fmt.Errorf("date must be between %s and %s", c.Min.Format(c.Format), c.Max.Format(c.Format))
	}

	return nil
}
