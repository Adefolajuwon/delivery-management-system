package main

import (
	"delivery-management-system/initializers"
	models "delivery-management-system/models"
)

func init() {
	initializers.ConnectDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Agent{}, &models.Warehouse{}, &models.AgentActivityLog{}, &models.Order{})
}
