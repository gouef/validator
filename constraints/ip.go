package constraints

import (
	"errors"
	"net"
)

type IsIPv4 struct{}

// Validate function for validate value
//
// Example:
//
//	con := constraints.IsIPv4{}
//	errs := validator.Validate(value, con)
func (c IsIPv4) Validate(value any) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("this value should be string type")
	}

	ip := net.ParseIP(str)
	if ip == nil || ip.To4() == nil {
		return errors.New("this value should be a valid IPv4 address")
	}

	return nil
}

type IsIPv6 struct{}

// Validate function for validate value
//
// Example:
//
//	con := constraints.IsIPv6{}
//	errs := validator.Validate(value, con)
func (c IsIPv6) Validate(value any) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("this value should be string type")
	}

	ip := net.ParseIP(str)
	if ip == nil || ip.To4() != nil {
		return errors.New("this value should be a valid IPv6 address")
	}

	return nil
}

type IsMAC struct{}

// Validate function for validate value
//
// Example:
//
//	con := constraints.IsMAC{}
//	errs := validator.Validate(value, con)
func (c IsMAC) Validate(value any) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("this value should be string type")
	}

	_, err := net.ParseMAC(str)
	if err != nil {
		return errors.New("this value should be a valid MAC address")
	}

	return nil
}
