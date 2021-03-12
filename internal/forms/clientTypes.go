package forms

import (
	"github.com/yvv4git/erp-fglaw/internal/domain"
	"gorm.io/gorm"
)

// ClientTypes form.
type ClientTypes struct {
	pagination Pagination
	ID         int64  `valid:"type(int64)" json:"id"`
	ClientType string `valid:"length(0|20)" json:"ctype"`
	ActingAs   string `valid:"length(0|20)" json:"actas"`
}

// ReadFirstPage is used to get a list of client-types entities on first page.
func (f *ClientTypes) ReadFirstPage(db *gorm.DB) (result []domain.ClientTypes, err error) {
	// Fill from form.
	var clientTypes domain.ClientTypes

	// Find in data storage.
	err = db.Where(&clientTypes).
		Limit(int(f.pagination.Limit)).
		Find(&result).Error

	return result, err
}

// ReadPage is used to find list of client-types entities by page number.
func (f *ClientTypes) ReadPage(db *gorm.DB) (result []domain.ClientTypes, err error) {
	// Fill from form.
	var clientTypes domain.ClientTypes

	// Find in data storage.
	err = db.Where(&clientTypes).
		Offset(f.pagination.Offset()).
		Limit(int(f.pagination.Limit)).
		Find(&result).Error

	return result, err
}

// Create is used for create client-types entity.
func (f *ClientTypes) Create(db *gorm.DB) (err error) {
	var clientTypes domain.ClientTypes

	clientTypes.ID = int(f.ID)
	clientTypes.ClientType = f.ClientType
	clientTypes.ActingAs = f.ActingAs

	return db.Create(&clientTypes).Error
}

// Update is used to update entity in data storage,
func (f *ClientTypes) Update(db *gorm.DB) (err error) {
	var clientTypes domain.ClientTypes

	// Find record in db.
	clientTypes.ID = int(f.ID)
	err = db.First(&clientTypes).Error
	if err != nil {
		return err
	}

	clientTypes.ClientType = f.ClientType
	clientTypes.ActingAs = f.ActingAs

	return db.Save(&clientTypes).Error
}

// Delete is used for delete entity from data storage.
func (f *ClientTypes) Delete(db *gorm.DB) (err error) {
	var clientTypes domain.ClientTypes

	clientTypes.ID = int(f.ID)
	return db.Delete(&clientTypes).Error
}
