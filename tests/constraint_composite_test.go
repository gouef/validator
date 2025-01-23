package tests

import (
	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNestedWithCompositeConstraint(t *testing.T) {
	nestedValidator := constraints.Field{
		Constraints: map[string]validator.Constraint{
			"Email": constraints.Composite{
				Constraints: []validator.Constraint{
					constraints.NotBlank{},
					constraints.Length{Min: 5, Max: 50},
					constraints.Email{},
				},
			},
			"Name": constraints.Composite{
				Constraints: []validator.Constraint{
					constraints.NotBlank{},
					constraints.Length{Min: 3, Max: 100},
				},
			},
		},
	}

	tests := []struct {
		name      string
		value     any
		expectErr bool
	}{
		{
			name: "Valid data",
			value: map[string]any{
				"Email": "test@example.com",
				"Name":  "John Doe",
			},
			expectErr: false,
		},
		{
			name: "Invalid Email",
			value: map[string]any{
				"Email": "invalid",
				"Name":  "John Doe",
			},
			expectErr: true,
		},
		{
			name: "Too short Name",
			value: map[string]any{
				"Email": "test@example.com",
				"Name":  "JD",
			},
			expectErr: true,
		},
		{
			name: "Blank fields",
			value: map[string]any{
				"Email": "",
				"Name":  "",
			},
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := nestedValidator.Validate(test.value)
			if test.expectErr {
				assert.Error(t, err, "Expected an error but got none")
			} else {
				assert.NoError(t, err, "Expected no error but got one")
			}
		})
	}
}
