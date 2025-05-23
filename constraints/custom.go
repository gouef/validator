package constraints

import "errors"

// Custom allows validating a value using a user-defined function.
// If the function returns an error, the constraint fails.
type Custom struct {
	Fn func(value any) error
}

func (c Custom) Validate(value any) error {
	if c.Fn == nil {
		return errors.New("custom constraint: missing function")
	}
	return c.Fn(value)
}
