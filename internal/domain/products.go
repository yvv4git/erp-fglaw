package domain

// Products - model for db entity.
type Products struct {
	ID             int `gorm:"primaryKey;autoIncrement:true"`
	SuppliersID    int
	ProductTypeID  int
	ProductName    string
	ProductBrand   string
	ProductPhotoID int
	Supplier       Suppliers    `gorm:"foreignKey:SuppliersID;references:ID"`
	ProductType    ProductTypes `gorm:"foreignKey:ProductTypeID;references:ID"`
	ProductPhoto   Photos       `gorm:"foreignKey:ProductPhotoID;references:ID"`
}
