<img align=right width="168" src="docs/gouef_logo.png">

# Validator
This package provides tools for validating various types of data in Go using predefined and user-defined validators.

[![GoDoc](https://pkg.go.dev/badge/github.com/gouef/validator.svg)](https://pkg.go.dev/github.com/gouef/validator)
[![GitHub stars](https://img.shields.io/github/stars/gouef/validator?style=social)](https://github.com/gouef/validator/stargazers)
[![Go Report Card](https://goreportcard.com/badge/github.com/gouef/validator)](https://goreportcard.com/report/github.com/gouef/validator)
[![codecov](https://codecov.io/github/gouef/validator/branch/main/graph/badge.svg?token=YUG8EMH6Q8)](https://codecov.io/github/gouef/validator)

## Versions
![Stable Version](https://img.shields.io/github/v/release/gouef/validator?label=Stable&labelColor=green)
![GitHub Release](https://img.shields.io/github/v/release/gouef/validator?label=RC&include_prereleases&filter=*rc*&logoSize=diago)
![GitHub Release](https://img.shields.io/github/v/release/gouef/validator?label=Beta&include_prereleases&filter=*beta*&logoSize=diago)

## Installation

If you're using Go modules, add this package to your project with the following command:

```bash
go get -u github.com/gouef/validator
```

## Usages

### Basic Validation
This library allows easy validation of values using defined `Constraint` types.

```go
package main

import (
	"fmt"
	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
)

func main() {
	// Example usage for email validation
	emailValidator := constraints.Email{}
	errs := validator.Validate("example@example.com", emailValidator)

	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Println("Error: ", err)
		}
	} else {
		fmt.Println("Email is valid.")
	}

	// Example validation for a value within a range
	rangeValidator := constraints.RangeConstraint{Min: 10, Max: 100}
	errs = validator.Validate(50, rangeValidator)

	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Println("Error: ", err)
		}
	} else {
		fmt.Println("Value is within the allowed range.")
	}
	
	// Example multiple constrains
	errs := validator.Validate("test@example.com", constraints.NotBlank{}, constraints.Email{})
	
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Println("Error: ", err)
		}
	} else {
		fmt.Println("Value is valid.")
	}
}

```

### Custom Validator
If you want to add your own validator, simply implement the `Constraint` interface and write the `Validate` method to perform the validation.

```go
package constraints

import "errors"

type GreaterThan100 struct{}

func (c GreaterThan100) Validate(value any) error {
	num, ok := value.(int)
	if !ok {
		return errors.New("this value must be an integer")
	}
	if num <= 100 {
		return errors.New("this value must be greater than 100")
	}
	return nil
}

```

You can use this custom validator just like the others:

```go
gtValidator := constraints.GreaterThan100{}
errs := validator.Validate(150, gtValidator)

```

Or create [Issue](/issues) or [PR](/pulls).


### Available Validators

This package includes several predefined validators:

- `Blank`: Ensures the value is blank (`""` or `nil`).
- `NotBlank`: Ensures the value is not blank.
- `IsTrue`: Ensures the value is `true`.
- `IsFalse`: Ensures the value is `false`.
- `IsBool`: Ensures the value is a boolean (`true` or `false`).
- `Currency`: Ensures the value is a valid currency code.
- `Email`: Ensures the value is a valid email address.
- `Language`: Ensures the value is a valid language code without a region (e.g., `en`).
- `LanguageWithRegion`: Ensures the value is a valid language code with a region (e.g., `en-US`).
- `Range`: Ensures the value is within a defined range (e.g., a number between `min` and `max`).
- `Required`: Ensures the value is not blank or `nil`.

More at [documentation](./docs/Validators.md)

### Configuration Options

Some validators allow configuration, for example:

- The `Email` validator has an `AllowNil()` method to specify whether `nil` is allowed as a valid value.

```go
emailValidator := constraints.Email{}
emailValidator.AllowNil() // Allows nil as a valid value
```

### Testing

You can test validators using the `testing` package:

```go
package validator_test

import (
	"testing"
	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
	"github.com/stretchr/testify/assert"
)

func TestEmailValidator(t *testing.T) {
	emailValidator := constraints.Email{}
	errs := validator.Validate("test@example.com", emailValidator)
	assert.Empty(t, errs)

	errs = validator.Validate("invalid-email", emailValidator)
	assert.NotEmpty(t, errs)
}
```

### Error Messages

When a validator returns an error, it is formatted as follows:

```go
errors.New("this value is required")
```

If the validation is successful, an empty slice of errors is returned.


## Contributing

Read [Contributing](CONTRIBUTING.md)

## Contributors

<div>
<span>
  <a href="https://github.com/JanGalek"><img src="https://raw.githubusercontent.com/gouef/validator/refs/heads/contributors-svg/.github/contributors/JanGalek.svg" alt="JanGalek" /></a>
</span>
<span>
  <a href="https://github.com/actions-user"><img src="https://raw.githubusercontent.com/gouef/validator/refs/heads/contributors-svg/.github/contributors/actions-user.svg" alt="actions-user" /></a>
</span>
</div>

