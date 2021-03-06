package models

// FileTypes - model for db entity.
type FileTypes struct {
	ID       int `gorm:"primaryKey;autoIncrement:true"`
	TypeName string
}
