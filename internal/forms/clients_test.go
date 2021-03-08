package forms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClients_Validate(t *testing.T) {
	testCases := []struct {
		name        string
		form        Clients
		result      bool
		wantErr     bool
		description string
	}{
		{
			name: "Good form",
			form: Clients{
				ID:           1,
				Number:       1,
				Address:      "San Francisco 88",
				CuitCustomer: "12345",
				ClientPhone:  "123456789012",
				ClientTypeID: 1,
			},
			result:      true,
			wantErr:     false,
			description: "This is good form",
		},
		{
			name: "Bad form",
			form: Clients{
				ID:           1,
				Number:       1,
				Address:      "123456789012345678901234567890123456789012345678901234567890",
				CuitCustomer: "12345",
				ClientPhone:  "123456789012",
				ClientTypeID: 1,
			},
			result:      false,
			wantErr:     true,
			description: "Very long address value",
		},
		{
			name: "Bad form-2",
			form: Clients{
				ID:           1,
				Number:       1,
				Address:      "12345678901234567890",
				CuitCustomer: "12345",
				ClientPhone:  "1234567890123456789012345",
				ClientTypeID: 1,
			},
			result:      false,
			wantErr:     true,
			description: "Very long client phone value",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Validate(tc.form)
			if tc.wantErr {
				assert.False(t, result, tc.description)
				assert.NotEmpty(t, err, tc.description)
			} else {
				assert.Equal(t, tc.result, result, tc.description)
			}
		})
	}
}
