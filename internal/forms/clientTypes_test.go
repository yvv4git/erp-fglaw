package forms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientTypes_Validate(t *testing.T) {
	testCases := []struct {
		name        string
		form        ClientTypes
		result      bool
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
			result:      true,
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
			result:      true,
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
			result:      true,
			wantErr:     true,
			description: "Very long words",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.form.Validate(tc.form)
			if tc.wantErr {
				assert.False(t, result, tc.description)
				assert.NotEmpty(t, tc.form.GetErrors(), tc.description)
			} else {
				assert.Equal(t, tc.result, result, tc.description)
			}
		})
	}
}
