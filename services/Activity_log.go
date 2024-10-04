package services

import (
	"delivery-management-system/dtos"
	"delivery-management-system/initializers"
	"delivery-management-system/models"
	"log"
	"time"
)

/*Create a new entity for each each agent */
func InitializeAgentActivity(agents []models.Agent) error {
	for _, agent := range agents {
		initialLog := models.AgentActivityLog{
			AgentID:           agent.AgentID,
			LogDate:           time.Now(),
			WarehouseID:       agent.WarehouseID,
			TransactionStatus: "active",
			TransactionType:   "shift_start",
		}

		if err := initializers.DB.Create(&initialLog).Error; err != nil {
			log.Println("Error creating initial log for agent:", agent.AgentID, err)
			return err
		}
	}

	return nil
}
