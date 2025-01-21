package constraints

import (
	"errors"
	"fmt"
	"golang.org/x/text/language"
	"regexp"
)

type LanguageWithRegion struct{}

type Language struct{}

func (c LanguageWithRegion) Validate(value any) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("this value must be a string")
	}

	matched, err := regexp.MatchString(`^[a-zA-Z]{2,3}-[a-zA-Z]{2,3}$`, str)

	if !matched || err != nil {
		return fmt.Errorf("invalid language-region format: %s", str)
	}

	tag, err := language.Parse(str)
	if err != nil {
		return fmt.Errorf("invalid language: %s", str)
	}

	region, _ := tag.Region()

	if !region.IsCountry() {
		return fmt.Errorf("invalid language code: %s (missing region)", str)
	}

	return nil
}

func (c Language) Validate(value any) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("this value must be a string")
	}

	matched, err := regexp.MatchString(`^[a-zA-Z]{2,3}$`, str)

	if !matched || err != nil {
		return fmt.Errorf("invalid language-region format: %s", str)
	}

	_, err = language.Parse(str)
	if err != nil {
		return fmt.Errorf("invalid language code: %s", str)
	}

	return nil
}
