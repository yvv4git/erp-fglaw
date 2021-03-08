package models

// Deposites - model for db entity.
type Deposites struct {
	ID          int `gorm:"primaryKey;autoIncrement:true"`
	Description string
}
