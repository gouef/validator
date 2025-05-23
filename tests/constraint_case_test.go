package tests

import (
	"testing"

	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
)

func TestIsLowercase(t *testing.T) {
	errs := validator.Validate("lowercase", constraints.IsLowercase{})
	if len(errs) != 0 {
		t.Errorf("unexpected error: %v", errs)
	}
	errs = validator.Validate("NotLower", constraints.IsLowercase{})
	if len(errs) == 0 {
		t.Errorf("expected error but got none")
	}
	errs = validator.Validate(12, constraints.IsLowercase{})
	if len(errs) == 0 {
		t.Errorf("expected error but got none")
	}
}

func TestIsUppercase(t *testing.T) {
	errs := validator.Validate("UPPERCASE", constraints.IsUppercase{})
	if len(errs) != 0 {
		t.Errorf("unexpected error: %v", errs)
	}
	errs = validator.Validate("Mixed", constraints.IsUppercase{})
	if len(errs) == 0 {
		t.Errorf("expected error but got none")
	}
	errs = validator.Validate(12, constraints.IsUppercase{})
	if len(errs) == 0 {
		t.Errorf("expected error but got none")
	}
}

func TestIsTitleCase(t *testing.T) {
	errs := validator.Validate("Title", constraints.IsTitleCase{})
	if len(errs) != 0 {
		t.Errorf("unexpected error: %v", errs)
	}
	errs = validator.Validate("title", constraints.IsTitleCase{})
	if len(errs) == 0 {
		t.Errorf("expected error but got none")
	}
	errs = validator.Validate(12, constraints.IsTitleCase{})
	if len(errs) == 0 {
		t.Errorf("expected error but got none")
	}
}

func TestIsCamelCase(t *testing.T) {
	errs := validator.Validate("myVariableName", constraints.IsCamelCase{})
	if len(errs) != 0 {
		t.Errorf("unexpected error: %v", errs)
	}
	errs = validator.Validate("MyVariable", constraints.IsCamelCase{})
	if len(errs) == 0 {
		t.Errorf("expected error but got none")
	}
	errs = validator.Validate(12, constraints.IsCamelCase{})
	if len(errs) == 0 {
		t.Errorf("expected error but got none")
	}
	errs = validator.Validate("myVariable Name", constraints.IsCamelCase{})
	if len(errs) == 0 {
		t.Errorf("expected error but got none")
	}
}

func TestIsPascalCase(t *testing.T) {
	errs := validator.Validate("MyVariableName", constraints.IsPascalCase{})
	if len(errs) != 0 {
		t.Errorf("unexpected error: %v", errs)
	}
	errs = validator.Validate("myVariable", constraints.IsPascalCase{})
	if len(errs) == 0 {
		t.Errorf("expected error but got none")
	}
	errs = validator.Validate(12, constraints.IsPascalCase{})
	if len(errs) == 0 {
		t.Errorf("expected error but got none")
	}
	errs = validator.Validate("MyVariable Name", constraints.IsPascalCase{})
	if len(errs) == 0 {
		t.Errorf("expected error but got none")
	}
}
