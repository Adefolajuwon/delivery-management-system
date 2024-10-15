package main

import (
	"delivery-management-system/initializers"
	models "delivery-management-system/models"
)

func init() {
	initializers.ConnectDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.AgentActivityLog{}, &models.Agent{}, &models.Order{}, &models.Warehouse{})
}
