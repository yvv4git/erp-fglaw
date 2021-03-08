package models

// ProductTypes - model for db entity.
type ProductTypes struct {
	ID       int `gorm:"primaryKey;autoIncrement:true"`
	TypeName string
}
