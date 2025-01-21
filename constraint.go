package validator

type Constraint interface {
	Validate(value any) error
}
