package models

type Warehouse struct {
	WarehouseID int     `gorm:"primaryKey;autoIncrement"`
	Name        string  `gorm:"type:varchar(255);not null"`
	Location    string
}
