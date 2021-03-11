package forms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientTypes_Validate(t *testing.T) {
	testCases := []struct {
		name        string
		form        ClientTypes
		wantErr     bool
		description string
	}{
		{
			name: "Good form",
			form: ClientTypes{
				ID:         1,
				ClientType: "Some",
				ActingAs:   "WTF some",
			},
			wantErr:     false,
			description: "This is good form",
		},
		{
			name: "Good form-2",
			form: ClientTypes{
				ID:         1,
				ClientType: "",
				ActingAs:   "",
			},
			wantErr:     false,
			description: "This is good form #2",
		},
		{
			name: "Bad form",
			form: ClientTypes{
				ID:         1,
				ClientType: "1234567890123456789012345",
				ActingAs:   "1234567890123456789012345",
			},
			wantErr:     true,
			description: "Very long words",
		},
		{
			name: "Pagination",
			form: ClientTypes{
				pagination: Pagination{
					Page:  1,
					Limit: 10,
				},
			},
			wantErr:     false,
			description: "Validate pagination",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := Validate(tc.form)
			if tc.wantErr {
				assert.NotEmpty(t, err, tc.description)
			} else {
				assert.Empty(t, err, tc.description)
			}
		})
	}
}
