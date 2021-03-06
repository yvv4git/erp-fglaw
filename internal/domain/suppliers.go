package models

// Suppliers is struct of model(поставщики).
type Suppliers struct {
	ID            int `gorm:"primaryKey;autoIncrement:true"`
	CuitSuppliers string
	Name          string
	Address       string
	Phone         string
}
