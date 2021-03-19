package forms

import (
	"github.com/yvv4git/erp-fglaw/internal/domain"
	"gorm.io/gorm"
)

// Clients ...
type Clients struct {
	Pagination   Pagination
	ID           int64  `valid:"type(int64)" json:"id"`
	Number       int64  `valid:"type(int64)" json:"num"`
	Address      string `valid:"length(0|50)" json:"addr"`
	CuitCustomer string `valid:"length(0|20)" json:"cuit"`
	ClientPhone  string `valid:"length(0|20)" json:"phone"`
	ClientTypeID int64  `valid:"type(int64)" json:"typeid"`
}

// ReadPage is used for find clients entities.
func (f *Clients) ReadPage(db *gorm.DB) (result []domain.Clients, err error) {
	// Fill from form.
	var clients domain.Clients
	clients.ID = int(f.ID)
	clients.Number = int(f.Number)
	clients.Address = f.Address
	clients.CuitCustomer = f.CuitCustomer
	clients.ClientPhone = f.ClientPhone
	clients.ClientTypeID = int(f.ClientTypeID)

	// Find in data storage.
	err = db.Where(&clients).
		Offset(f.Pagination.Offset()).
		Limit(int(f.Pagination.Limit)).
		Find(&result).Error

	return result, err
}

// Create is used for create clients entity in data storage.
func (f *Clients) Create(db *gorm.DB) (err error) {
	// Fill from form.
	var clients domain.Clients
	clients.ID = int(f.ID)
	clients.Number = int(f.Number)
	clients.Address = f.Address
	clients.CuitCustomer = f.CuitCustomer
	clients.ClientPhone = f.ClientPhone
	clients.ClientTypeID = int(f.ClientTypeID)

	return db.Create(&clients).Error
}

// Update is used for update entity in data storage.
func (f *Clients) Update(db *gorm.DB) (err error) {
	var clients domain.Clients

	// Find record in db.
	clients.ID = int(f.ID)
	db.First(&clients)

	// Fill from form.
	clients.ID = int(f.ID)
	clients.Number = int(f.Number)
	clients.Address = f.Address
	clients.CuitCustomer = f.CuitCustomer
	clients.ClientPhone = f.ClientPhone
	clients.ClientTypeID = int(f.ClientTypeID)

	return db.Save(&clients).Error
}

// Delete is used for delete clients entity from data storage.
func (f *Clients) Delete(db *gorm.DB) (err error) {
	var clients domain.Clients
	clients.ID = int(f.ID)
	return db.Delete(&clients).Error
}

// Count is used for provides the number of client entities relevant to the query.
func (f *Clients) Count(db *gorm.DB) (result int64, err error) {
	// Fill from form.
	var clients domain.Clients
	clients.ID = int(f.ID)
	clients.Number = int(f.Number)
	clients.Address = f.Address
	clients.CuitCustomer = f.CuitCustomer
	clients.ClientPhone = f.ClientPhone
	clients.ClientTypeID = int(f.ClientTypeID)

	// Get count of entities.
	err = db.Model(&domain.Clients{}).Where(&clients).Count(&result).Error
	return
}
