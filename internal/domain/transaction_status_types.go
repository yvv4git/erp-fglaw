package models

/*
TransactionStatusTypes - model of transaction types.
The status field should have only three or four options and this are:
1. DEPOSIT (means that the product is in my office storage).
2. DELIVERED (means that the product is already sent to a Hospital).
3. SOLD (when the product has been sold to either the hospital or insurance company and the invoice sell has been issued).
*/
type TransactionStatusTypes struct {
	ID   int `gorm:"primaryKey;autoIncrement:true"`
	Name string
}
