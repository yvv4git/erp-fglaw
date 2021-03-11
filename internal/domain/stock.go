package domain

// Stock is struct of model(склад).
type Stock struct {
	ID              int `gorm:"primaryKey;autoIncrement:true"`
	SeriesNumber    string
	ProductID       int
	ExpirationDate  string
	EntryDate       string
	InvoicePurchase int
	InvoiceSell     int
	TransactionID   int
}
