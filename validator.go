package validator

func Validate(value any, constraints ...Constraint) []error {
	var errs []error
	for _, constraint := range constraints {
		if err := constraint.Validate(value); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}
