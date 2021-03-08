package models

// ClientTypes - entity for db table.
type ClientTypes struct {
	ID         int `gorm:"primaryKey;autoIncrement:true"`
	ClientType string
	ActingAs   string
}
