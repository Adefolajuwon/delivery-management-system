package models

import "time"

type AgentActivityLog struct {
    LogID              uint      `gorm:"primaryKey;autoIncrement"`
    AgentID            string    `gorm:"type:uuid;not null"` // Foreign key to agents table
    LogDate            time.Time `gorm:"type:date;not null"` // Date for the log
    OrderID            uint      `gorm:"not null"` // Foreign key to orders table, nullable if no order was completed
    WarehouseID        int       `gorm:"not null"` // Foreign key to warehouses
    TotalHoursWorked   float64   `gorm:"type:decimal(5,2);default:0"` // Total hours worked for the day
    TotalDistanceCovered float64  `gorm:"type:decimal(6,2);default:0"` // Distance covered for the day
    TotalOrdersCompleted int      `gorm:"type:int;default:0"` // Total orders completed for the day
    TransactionTime    time.Time `gorm:"autoCreateTime"` // Timestamp for when the transaction occurred
    TransactionType    string    `gorm:"type:varchar(50);not null"` // e.g., "delivery", "return"
    TransactionStatus  string    `gorm:"type:varchar(50);not null"` // e.g., "completed", "failed"
    DistanceCoveredKM  float64   `gorm:"type:decimal(5,2);default:0"` // Distance for the transaction
    TimeSpentHours     float64   `gorm:"type:decimal(4,2);default:0"` // Time taken for the transaction
    PaymentAmount      float64   `gorm:"type:decimal(10,2);default:0"` // Payment amount if applicable
    Notes              string    `gorm:"type:text"` // Optional notes for the transaction
}
