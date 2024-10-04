package models

import "time"

type AgentActivityLog struct {
	LogID                int       `gorm:"primaryKey;autoIncrement"`
	AgentID              int       `gorm:"not null;index"`           // Foreign key reference to Agent
	LogDate              time.Time `gorm:"type:date;not null"`
	OrderID              *int      `gorm:"default:null"`             // Nullable foreign key reference to Order
	WarehouseID          int       `gorm:"not null;index"`           // Foreign key reference to Warehouse
	TotalHoursWorked     float64   `gorm:"type:decimal(5,2);default:0"`
	TotalDistanceCovered float64   `gorm:"type:decimal(6,2);default:0"`
	TotalOrdersCompleted int       `gorm:"default:0"`
	TransactionTime      time.Time `gorm:"autoCreateTime"`
	TransactionType      string    `gorm:"type:varchar(50);not null"`
	TransactionStatus    string    `gorm:"type:varchar(50);not null"`
	DistanceCoveredKM    float64   `gorm:"type:decimal(5,2);default:0"`
	TimeSpentHours       float64   `gorm:"type:decimal(4,2);default:0"`
	PaymentAmount        float64   `gorm:"type:decimal(10,2);default:0"`
	Notes                string    `gorm:"type:text"`
}