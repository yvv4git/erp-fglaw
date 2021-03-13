package forms_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yvv4git/erp-fglaw/internal/domain"
	"github.com/yvv4git/erp-fglaw/internal/forms"
	"github.com/yvv4git/erp-fglaw/tests"
	"gorm.io/gorm"
)

func TestClientTypes_Validate(t *testing.T) {
	testCases := []struct {
		name        string
		form        forms.ClientTypes
		wantErr     bool
		description string
	}{
		{
			name: "Good form",
			form: forms.ClientTypes{
				ID:         1,
				ClientType: "Some",
				ActingAs:   "WTF some",
			},
			wantErr:     false,
			description: "This is good form",
		},
		{
			name: "Good form-2",
			form: forms.ClientTypes{
				ID:         1,
				ClientType: "",
				ActingAs:   "",
			},
			wantErr:     false,
			description: "This is good form #2",
		},
		{
			name: "Bad form",
			form: forms.ClientTypes{
				ID:         1,
				ClientType: "1234567890123456789012345",
				ActingAs:   "1234567890123456789012345",
			},
			wantErr:     true,
			description: "Very long words",
		},
		{
			name: "Pagination",
			form: forms.ClientTypes{
				Pagination: forms.Pagination{
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
			err := forms.Validate(tc.form)
			if tc.wantErr {
				assert.NotEmpty(t, err, tc.description)
			} else {
				assert.Empty(t, err, tc.description)
			}
		})
	}
}

func TestClientTypes_ReadFirstPage(t *testing.T) {
	testCases := []struct {
		name        string
		form        forms.ClientTypes
		wantErr     bool
		description string
	}{
		{
			name:        "Read first page",
			form:        forms.ClientTypes{},
			wantErr:     false,
			description: "Find data entity in first page",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tests.PrepareTestDatabase(fixtures)

			result, err := tc.form.ReadFirstPage(db)
			if tc.wantErr {
				assert.NotEmpty(t, err, tc.description)
			} else {
				assert.NotEmpty(t, result, tc.description)
				assert.Empty(t, err, tc.description)
			}
			//t.Log(result)
		})
	}
}

func TestClientTypes_ReadPage(t *testing.T) {
	testCases := []struct {
		name        string
		form        forms.ClientTypes
		wantErr     bool
		description string
	}{
		{
			name: "Read first page",
			form: forms.ClientTypes{
				Pagination: forms.Pagination{
					Page:  0,
					Limit: 10,
				},
			},
			wantErr:     false,
			description: "Read list client-types entities from first page.",
		},
		{
			name: "Read second page",
			form: forms.ClientTypes{
				Pagination: forms.Pagination{
					Page:  1,
					Limit: 10,
				},
			},
			wantErr:     false,
			description: "Read list client-types entities from second page.",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tests.PrepareTestDatabase(fixtures)

			result, err := tc.form.ReadPage(db)
			if tc.wantErr {
				assert.NotEmpty(t, err, tc.description)
			} else {
				assert.NotEmpty(t, result, tc.description)
				assert.Empty(t, err, tc.description)
			}
			//t.Log(result)
		})
	}
}

func TestClientTypes_Create(t *testing.T) {
	testCases := []struct {
		name        string
		form        forms.ClientTypes
		wantErr     bool
		check       func(db *gorm.DB) bool
		description string
	}{
		{
			name: "Create new client-types",
			form: forms.ClientTypes{
				ClientType: "NewClientType-12",
				ActingAs:   "Some information about actiong as...12",
			},
			wantErr: false,
			check: func(db *gorm.DB) bool {
				var clientTypes domain.ClientTypes
				err := db.Where("client_type = ?", "NewClientType-12").First(&clientTypes).Error
				if err != nil {
					return false
				}
				return true
			},
			description: "Create client-entity in data storage and check it.",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tests.PrepareTestDatabase(fixtures)

			err := tc.form.Create(db)
			if tc.wantErr {
				assert.NotNil(t, err, tc.description)
			} else {
				assert.Nil(t, err, tc.description)
				status := tc.check(db)
				assert.True(t, status, "Entity success founded in data storage.")
			}
		})
	}
}

func TestClientTypes_Update(t *testing.T) {
	testCases := []struct {
		name        string
		form        forms.ClientTypes
		check       func(db *gorm.DB) bool
		wantErr     bool
		description string
	}{
		{
			name: "Update first entity",
			form: forms.ClientTypes{
				ID:         1,
				ClientType: "Orba Social",
				ActingAs:   "Updated value",
			},
			check: func(db *gorm.DB) bool {
				var clientTypes domain.ClientTypes
				err := db.First(&clientTypes, 10).Error
				if err != nil {
					return false
				}
				return true
			},
			wantErr:     false,
			description: "Update first entity in data storage. Change ActingAs.",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tests.PrepareTestDatabase(fixtures)

			err := tc.form.Update(db)
			if tc.wantErr {
				assert.NotNil(t, err, tc.description)
			} else {
				assert.Nil(t, err, tc.description)
				status := tc.check(db)
				assert.True(t, status, "Check entity was updated.")
			}
		})
	}
}

func TestClientTypes_Delete(t *testing.T) {
	testCases := []struct {
		name        string
		form        forms.ClientTypes
		check       func(db *gorm.DB) bool
		wantErr     bool
		description string
	}{
		{
			name: "Delete entity by id",
			form: forms.ClientTypes{
				ID: 1,
			},
			check: func(db *gorm.DB) bool {
				var clientTypes domain.ClientTypes
				clientTypes.ID = 1
				err := db.First(&clientTypes).Error
				if err != nil {
					return (err == gorm.ErrRecordNotFound)
				}
				return true
			},
			wantErr:     false,
			description: "Delete entity from data storage by id.",
		},
		{
			name: "Delete entity by id, but no found",
			form: forms.ClientTypes{
				ID: 100500,
			},
			check: func(db *gorm.DB) bool {
				var clientTypes domain.ClientTypes
				clientTypes.ID = 1
				err := db.First(&clientTypes).Error
				if err != nil {
					return (err == gorm.ErrRecordNotFound)
				}
				return true
			},
			wantErr:     false,
			description: "Delete entity from data storage by id, but not found",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tests.PrepareTestDatabase(fixtures)

			err := tc.form.Delete(db)
			if tc.wantErr {
				assert.NotNil(t, err, tc.description)
			} else {
				assert.Nil(t, err, tc.description)
				status := tc.check(db)
				assert.True(t, status, "Check entity was updated.")
			}
		})
	}
}
