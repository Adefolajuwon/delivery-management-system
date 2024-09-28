package models

type Warehouse struct {
	WarehouseID  int     `gorm:"primaryKey;autoIncrement"`
	Name         string  `gorm:"type:varchar(255);not null"`
	LocationLat  float64 `gorm:"type:decimal(10,8);not null"`
	LocationLong float64 `gorm:"type:decimal(11,8);not null"`
}
