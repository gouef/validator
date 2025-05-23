package tests

import (
	"testing"

	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
)

func TestIsUUID(t *testing.T) {
	t.Run("valid UUID (generic)", func(t *testing.T) {
		errs := validator.Validate("550e8400-e29b-41d4-a716-446655440000", constraints.IsUUID{})
		if len(errs) != 0 {
			t.Errorf("unexpected errors: %v", errs)
		}
	})

	t.Run("invalid UUID format", func(t *testing.T) {
		errs := validator.Validate("not-a-uuid", constraints.IsUUID{})
		if len(errs) == 0 {
			t.Errorf("expected error but got none")
		}
	})

	t.Run("valid UUID v4", func(t *testing.T) {
		errs := validator.Validate("f47ac10b-58cc-4372-a567-0e02b2c3d479", constraints.IsUUID{Version: 4})
		if len(errs) != 0 {
			t.Errorf("unexpected error: %v", errs)
		}
	})

	t.Run("wrong UUID version", func(t *testing.T) {
		errs := validator.Validate("f47ac10b-58cc-1372-a567-0e02b2c3d479", constraints.IsUUID{Version: 4})
		if len(errs) == 0 {
			t.Errorf("expected error for wrong version but got none")
		}
	})

	t.Run("unsupported UUID version (2)", func(t *testing.T) {
		errs := validator.Validate("550e8400-e29b-41d4-a716-446655440000", constraints.IsUUID{Version: 2})
		if len(errs) == 0 {
			t.Errorf("expected error for unsupported version but got none")
		} else if errs[0].Error() != "unsupported UUID version: 2" {
			t.Errorf("unexpected error: %v", errs[0])
		}
	})
}

func TestIsUUIDVersions(t *testing.T) {
	cases := []struct {
		name       string
		value      any
		constraint validator.Constraint
		valid      bool
	}{
		{"v1", "6ba7b810-9dad-11d1-80b4-00c04fd430c8", constraints.IsUUIDv1{}, true},
		{"v3", "f47ac10b-58cc-3372-a567-0e02b2c3d479", constraints.IsUUIDv3{}, true},
		{"v4", "f47ac10b-58cc-4372-a567-0e02b2c3d479", constraints.IsUUIDv4{}, true},
		{"v5", "f47ac10b-58cc-5372-a567-0e02b2c3d479", constraints.IsUUIDv5{}, true},
		{"v2 (invalid)", "f47ac10b-58cc-2372-a567-0e02b2c3d479", constraints.IsUUIDv2{}, false},
		{"invalid value type", 200, constraints.IsUUID{}, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			errs := validator.Validate(tc.value, tc.constraint)
			if tc.valid && len(errs) > 0 {
				t.Errorf("unexpected error: %v", errs)
			}
			if !tc.valid && len(errs) == 0 {
				t.Errorf("expected error but got none")
			}
		})
	}
}
