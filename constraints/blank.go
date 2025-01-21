package constraints

import "errors"

type Blank struct{}
type NotBlank struct{}

func (c Blank) Validate(value any) error {
	if value == "" || value == nil {
		return nil
	}

	return errors.New("this value should be blank")
}

func (c NotBlank) Validate(value any) error {
	if value != "" && value != nil {
		return nil
	}

	return errors.New("this value should not be blank")
}
