package models

type Agent struct {
	AgentID     int    `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(255);not null"`
	Phone       string `gorm:"type:varchar(15);not null;unique"`
	WarehouseID int    `gorm:"not null;index;foreignKey:WarehouseID;references:WarehouseID"` // Foreign key to Warehouse
}
