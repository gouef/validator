package constraints

import (
	"errors"
	"fmt"
	"github.com/gouef/country"
)

type Country struct{}

// Validate function for validate value
//
// Example:
//
//	con := constraints.Country{}
//	value := "CZ"
//	errs := validator.Validate(value, con)
func (c Country) Validate(value any) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("this value must be a string")
	}

	alpha2 := country.FindByAlpha2(str)

	if alpha2 != nil {
		return nil
	}

	alpha3 := country.FindByAlpha3(str)

	if alpha3 != nil {
		return nil
	}

	name := country.FindByName(str)

	if name != nil {
		return nil
	}

	return fmt.Errorf("invalid country: %s", str)
}
