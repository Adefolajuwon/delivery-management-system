package models

import "time"

type Order struct {
	OrderID          int       `gorm:"primaryKey;autoIncrement"`
	WarehouseID      int       `gorm:"not null;index;foreignKey:WarehouseID;references:WarehouseID"` // Foreign key to Warehouse
	DestinationLat   float64   `gorm:"type:decimal(10,8);not null"`
	DestinationLong  float64   `gorm:"type:decimal(11,8);not null"`
	DeliveryAddress  string    `gorm:"type:text;not null"`
	AssignedAgentID  *int      `gorm:"default:null;foreignKey:AgentID;references:AgentID"` // Foreign key to Agent, can be null
	DeliveryStatus   string    `gorm:"type:varchar(50);default:'pending'"`
	OrderCreatedTime time.Time `gorm:"autoCreateTime"`
}
