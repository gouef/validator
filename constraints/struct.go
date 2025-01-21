package constraints

import (
	"errors"
	"fmt"
	"reflect"
)

type IsStruct struct{}

type Struct struct {
	Struct interface{}
}

func (c IsStruct) Validate(value any) error {
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Struct {
		return errors.New("this value must be a struct")
	}

	return nil
}

func (c Struct) Validate(value any) error {
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Struct {
		return errors.New("this value must be a struct")
	}

	expectedType := reflect.TypeOf(c.Struct)
	if expectedType.Kind() == reflect.Struct {
		for i := 0; i < expectedType.NumField(); i++ {
			fieldName := expectedType.Field(i).Name
			fieldValue := v.FieldByName(fieldName)

			if !fieldValue.IsValid() {
				return fmt.Errorf("struct does not contain field '%s'", fieldName)
			}
		}
	}

	return nil
}
