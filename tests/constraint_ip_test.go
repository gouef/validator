package tests

import (
	"testing"

	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
)

func TestIsIPv4(t *testing.T) {
	t.Run("valid IPv4", func(t *testing.T) {
		err := validator.Validate("192.168.1.1", constraints.IsIPv4{})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("invalid IPv4 format", func(t *testing.T) {
		err := validator.Validate("not-an-ip", constraints.IsIPv4{})
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("IPv6 input", func(t *testing.T) {
		err := validator.Validate("::1", constraints.IsIPv4{})
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("non-string input", func(t *testing.T) {
		err := validator.Validate(123, constraints.IsIPv4{})
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})
}

func TestIsIPv6(t *testing.T) {
	t.Run("valid IPv6", func(t *testing.T) {
		err := validator.Validate("2001:db8::1", constraints.IsIPv6{})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("IPv4 input", func(t *testing.T) {
		err := validator.Validate("192.168.1.1", constraints.IsIPv6{})
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("invalid input", func(t *testing.T) {
		err := validator.Validate("invalid-ip", constraints.IsIPv6{})
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("non-string input", func(t *testing.T) {
		err := validator.Validate(false, constraints.IsIPv6{})
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})
}

func TestIsMAC(t *testing.T) {
	t.Run("valid MAC", func(t *testing.T) {
		err := validator.Validate("00:1A:2B:3C:4D:5E", constraints.IsMAC{})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("valid MAC with dashes", func(t *testing.T) {
		err := validator.Validate("01-23-45-67-89-AB", constraints.IsMAC{})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("invalid MAC", func(t *testing.T) {
		err := validator.Validate("not-a-mac", constraints.IsMAC{})
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("non-string input", func(t *testing.T) {
		err := validator.Validate(12345, constraints.IsMAC{})
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})
}
