package models

type Agent struct {
	AgentID     string `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name        string `gorm:"type:varchar(255);not null"`
	Phone       string `gorm:"type:varchar(15);not null"`
	WarehouseID int
}
