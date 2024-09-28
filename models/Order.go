package models

import "time"

type Order struct {
	OrderID          uint      `gorm:"primaryKey;autoIncrement"`
	WarehouseID      int       `gorm:"not null"`
	DestinationLat   float64   `gorm:"type:decimal(10,8);not null"`
	DestinationLong  float64   `gorm:"type:decimal(11,8);not null"`
	DeliveryAddress  string    `gorm:"type:text;not null"`
	AssignedAgentID  string    `gorm:"type:uuid";default: 'null'`
	DeliveryStatus   string    `gorm:"type:varchar(50);default:'pending'"`
	OrderCreatedTime time.Time `gorm:"autoCreateTime"`
}
