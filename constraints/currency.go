package constraints

import (
	"errors"
	"fmt"
	"github.com/gouef/currency"
)

type Currency struct{}

func (c Currency) Validate(value any) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("this value must be a string")
	}

	valid := currency.ValidateCurrency(str)
	if valid == false {
		return fmt.Errorf("invalid language code: %s", str)
	}

	return nil
}
