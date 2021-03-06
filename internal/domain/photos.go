package models

// Photos - model for db entity.
type Photos struct {
	ID         int `gorm:"primaryKey;autoIncrement:true"`
	FileName   string
	FileTypeID int
	FileType   FileTypes `gorm:"foreignKey:FileTypeID;references:ID"`
}
