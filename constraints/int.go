package constraints

import "errors"

type IsInt struct{}
type IsInt8 struct{}
type IsInt16 struct{}
type IsInt32 struct{}
type IsInt64 struct{}

type IsUint struct{}
type IsUint8 struct{}
type IsUint16 struct{}
type IsUint32 struct{}
type IsUint64 struct{}

// Validate function for validate value
//
// Example:
//
//	con := constraints.IsInt{}
//	errs := validator.Validate(value, con)
func (c IsInt) Validate(value any) error {
	if _, ok := value.(int); ok {
		return nil
	}

	return errors.New("this value should be int type")
}

// Validate function for validate value
//
// Example:
//
//	con := constraints.IsInt8{}
//	errs := validator.Validate(value, con)
func (c IsInt8) Validate(value any) error {
	if _, ok := value.(int8); ok {
		return nil
	}

	return errors.New("this value should be int8 type")
}

// Validate function for validate value
//
// Example:
//
//	con := constraints.IsInt16{}
//	errs := validator.Validate(value, con)
func (c IsInt16) Validate(value any) error {
	if _, ok := value.(int16); ok {
		return nil
	}

	return errors.New("this value should be int16 type")
}

// Validate function for validate value
//
// Example:
//
//	con := constraints.IsInt32{}
//	errs := validator.Validate(value, con)
func (c IsInt32) Validate(value any) error {
	if _, ok := value.(int32); ok {
		return nil
	}

	return errors.New("this value should be int32 type")
}

// Validate function for validate value
//
// Example:
//
//	con := constraints.IsInt64{}
//	errs := validator.Validate(value, con)
func (c IsInt64) Validate(value any) error {
	if _, ok := value.(int64); ok {
		return nil
	}

	return errors.New("this value should be int64 type")
}

// Validate function for validate value
//
// Example:
//
//	con := constraints.IsUint{}
//	errs := validator.Validate(value, con)
func (c IsUint) Validate(value any) error {
	if _, ok := value.(uint); ok {
		return nil
	}

	return errors.New("this value should be uint type")
}

// Validate function for validate value
//
// Example:
//
//	con := constraints.IsUint8{}
//	errs := validator.Validate(value, con)
func (c IsUint8) Validate(value any) error {
	if _, ok := value.(uint8); ok {
		return nil
	}

	return errors.New("this value should be uint8 type")
}

// Validate function for validate value
//
// Example:
//
//	con := constraints.IsUint16{}
//	errs := validator.Validate(value, con)
func (c IsUint16) Validate(value any) error {
	if _, ok := value.(uint16); ok {
		return nil
	}

	return errors.New("this value should be uint16 type")
}

// Validate function for validate value
//
// Example:
//
//	con := constraints.IsUint32{}
//	errs := validator.Validate(value, con)
func (c IsUint32) Validate(value any) error {
	if _, ok := value.(uint32); ok {
		return nil
	}

	return errors.New("this value should be uint32 type")
}

// Validate function for validate value
//
// Example:
//
//	con := constraints.IsUint64{}
//	errs := validator.Validate(value, con)
func (c IsUint64) Validate(value any) error {
	if _, ok := value.(uint64); ok {
		return nil
	}

	return errors.New("this value should be uint64 type")
}
