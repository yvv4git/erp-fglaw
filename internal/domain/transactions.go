package models

// Transactions is struct of model.
type Transactions struct {
	ID              int `gorm:"primaryKey;autoIncrement:true"`
	TransactionDate string
	Status          string
	DeliveryDays    int
	SeriesNumber    string
	ClientID        int
	DepositeID      int
	ProductID       int
	Client          Clients   `gorm:"foreignKey:ClientID;references:ID"`
	Deposite        Deposites `gorm:"foreignKey:DepositeID;references:ID"`
	Product         Products  `gorm:"foreignKey:ProductID;references:ID"`
}
