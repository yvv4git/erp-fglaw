package models

// Clients - model for db entity.
type Clients struct {
	ID           int `gorm:"primaryKey;autoIncrement:true"`
	Number       int
	Address      string
	CuitCustomer string
	ClientPhone  string
	ClientTypeID int
	ClientType   ClientTypes `gorm:"foreignKey:ClientTypeID;references:ID"`
}
