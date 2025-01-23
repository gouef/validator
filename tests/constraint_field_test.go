package tests

import (
	"testing"

	"github.com/gouef/validator"
	"github.com/gouef/validator/constraints"
	"github.com/stretchr/testify/assert"
)

func TestField(t *testing.T) {
	nestedValidator := constraints.Field{
		Constraints: map[string]validator.Constraint{
			"Name":  constraints.NotBlank{},
			"Email": constraints.Email{},
			"Address": constraints.Field{
				Constraints: map[string]validator.Constraint{
					"City":    constraints.NotBlank{},
					"Country": constraints.NotBlank{},
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
				"Name":  "John Doe",
				"Email": "johndoe@example.com",
				"Address": map[string]any{
					"City":    "New York",
					"Country": "USA",
				},
			},
			expectErr: false,
		},
		{
			name: "Missing field in Address",
			value: map[string]any{
				"Name":  "John Doe",
				"Email": "johndoe@example.com",
				"Address": map[string]any{
					"City": "",
				},
			},
			expectErr: true,
		},
		{
			name: "Invalid Email",
			value: map[string]any{
				"Name":  "John Doe",
				"Email": "invalid-email",
				"Address": map[string]any{
					"City":    "New York",
					"Country": "USA",
				},
			},
			expectErr: true,
		},
		{
			name: "Empty Name",
			value: map[string]any{
				"Name":  "",
				"Email": "johndoe@example.com",
				"Address": map[string]any{
					"City":    "New York",
					"Country": "USA",
				},
			},
			expectErr: true,
		},
		{
			name:      "Nil value",
			value:     nil,
			expectErr: true,
		},
		{
			name:      "Invalid type for value",
			value:     "invalid-type",
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
