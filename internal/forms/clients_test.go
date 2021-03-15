package forms_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yvv4git/erp-fglaw/internal/domain"
	"github.com/yvv4git/erp-fglaw/internal/forms"
	"github.com/yvv4git/erp-fglaw/tests"
	"gorm.io/gorm"
)

func TestClients_Validate(t *testing.T) {
	testCases := []struct {
		name        string
		form        forms.Clients
		wantErr     bool
		description string
	}{
		{
			name: "Good form",
			form: forms.Clients{
				ID:           1,
				Number:       1,
				Address:      "San Francisco 88",
				CuitCustomer: "12345",
				ClientPhone:  "123456789012",
				ClientTypeID: 1,
			},
			wantErr:     false,
			description: "This is good form",
		},
		{
			name: "Bad form",
			form: forms.Clients{
				ID:           1,
				Number:       1,
				Address:      "123456789012345678901234567890123456789012345678901234567890",
				CuitCustomer: "12345",
				ClientPhone:  "123456789012",
				ClientTypeID: 1,
			},
			wantErr:     true,
			description: "Very long address value",
		},
		{
			name: "Bad form-2",
			form: forms.Clients{
				ID:           1,
				Number:       1,
				Address:      "12345678901234567890",
				CuitCustomer: "12345",
				ClientPhone:  "1234567890123456789012345",
				ClientTypeID: 1,
			},
			wantErr:     true,
			description: "Very long client phone value",
		},
		{
			name: "Pagination",
			form: forms.Clients{
				Pagination: forms.Pagination{
					Page:  1,
					Limit: 10,
				},
			},
			wantErr:     false,
			description: "Check pagination validate",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := forms.Validate(tc.form)
			if tc.wantErr {
				assert.NotEmpty(t, err, tc.description)
			} else {
				assert.Empty(t, err, tc.description)
			}
		})
	}
}

func TestClientsForm_ReadPage(t *testing.T) {
	testCases := []struct {
		name        string
		form        forms.Clients
		check       func(db *gorm.DB, result []domain.Clients)
		wantErr     bool
		description string
	}{
		{
			name: "Find first page.",
			form: forms.Clients{
				Pagination: forms.Pagination{
					Page:  0,
					Limit: 10,
				},
			},
			check: func(db *gorm.DB, result []domain.Clients) {
				assert.Equal(t, 3, len(result))
			},
			wantErr:     false,
			description: "Find clients entities for first page.",
		},
		{
			name: "Find second page.",
			form: forms.Clients{
				Pagination: forms.Pagination{
					Page:  1,
					Limit: 10,
				},
			},
			check: func(db *gorm.DB, result []domain.Clients) {
				assert.Equal(t, 0, len(result))
			},
			wantErr:     false,
			description: "Find clients entities for second page. No found.",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tests.PrepareTestDatabase()

			result, err := tc.form.ReadPage(db)
			if tc.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				tc.check(db, result)
			}
		})
	}
}

func TestClientsForm_Create(t *testing.T) {
	testCases := []struct {
		name        string
		form        forms.Clients
		check       func(db *gorm.DB)
		wantErr     bool
		description string
	}{
		{
			name: "Create clients entity.",
			form: forms.Clients{
				Number:       4,
				Address:      "2790 Green Street San-Francisco, CA 94123",
				CuitCustomer: "Some unknown string",
				ClientPhone:  "+17133373300",
				ClientTypeID: 1,
			},
			check: func(db *gorm.DB) {
				var clients domain.Clients
				clients.Number = 4
				err := db.
					Joins("JOIN client_types on client_types.id = clients.client_type_id").
					Preload("ClientType").
					Where("number = ?", 4).
					First(&clients).Error
				assert.Nil(t, err, "Find new client")
				assert.Equal(t, "Orba Social", clients.ClientType.ClientType, "Relation join table")
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tests.PrepareTestDatabase()

			err := tc.form.Create(db)
			if tc.wantErr {
				assert.NotNil(t, err, tc.description)
			} else {
				assert.Nil(t, err, tc.description)
				tc.check(db)
			}
		})
	}
}

func TestClientsForm_Update(t *testing.T) {
	testCases := []struct {
		name        string
		form        forms.Clients
		check       func(db *gorm.DB)
		wantErr     bool
		description string
	}{
		{
			name: "Update first client",
			form: forms.Clients{
				ID:           1,
				Number:       1,
				Address:      "1-2790 Green Street San-Francisco, CA 94123",
				CuitCustomer: "Some unknown string",
				ClientPhone:  "+17133373301",
				ClientTypeID: 1,
			},
			wantErr: false,
			check: func(db *gorm.DB) {
				var clients domain.Clients
				err := db.First(&clients, 1).Error
				assert.Nil(t, err)
				assert.Equal(t, "1-2790 Green Street San-Francisco, CA 94123", clients.Address)
			},
			description: "Update first client and then check it",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tests.PrepareTestDatabase()

			err := tc.form.Update(db)
			if tc.wantErr {
				assert.NotNil(t, err, tc.description)
			} else {
				assert.Nil(t, err, tc.description)
				tc.check(db)
			}
		})
	}
}

func TestClientsForm_Delete(t *testing.T) {
	testCases := []struct {
		name        string
		form        forms.Clients
		check       func(db *gorm.DB)
		wantErr     bool
		description string
	}{
		{
			name: "Delete third client",
			form: forms.Clients{
				ID: 3,
			},
			check: func(db *gorm.DB) {
				var clients domain.Clients
				err := db.First(&clients, 3).Error
				assert.NotNil(t, err)
				assert.Equal(t, err.Error(), gorm.ErrRecordNotFound.Error(), "This row not be found")
			},
			wantErr:     false,
			description: "Delete third client from data storage and check it",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tests.PrepareTestDatabase()

			err := tc.form.Delete(db)
			if tc.wantErr {
				assert.NotNil(t, err, tc.description)
			} else {
				assert.Nil(t, err, tc.description)
				tc.check(db)
			}
		})
	}
}
