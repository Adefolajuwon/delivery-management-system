package services

import (
	"delivery-management-system/dtos"
	"delivery-management-system/initializers"
	"delivery-management-system/models"
	"log"
	"delivery-management-system/helper"
)

func AllocateOrders() dtos.Response {
	var agents []models.Agent
	if err := initializers.DB.Find(&agents).Error; err != nil {
		log.Println("Error fetching agents:", err)
	}

	var orders []models.Order
	if err := initializers.DB.Where("assigned_agent_id IS NULL").Find(&orders).Error; err != nil {
		log.Println("Error fetching orders:", err)
	}

	// Fetch warehouses from the correct model
	var warehouses []models.Warehouse
	if err := initializers.DB.Find(&warehouses).Error; err != nil {
		log.Println("Error fetching warehouses:", err)
	}

	// Create a map to store agents grouped by their warehouse
	agentsInWarehouse := make(map[int][]models.Agent)
	// Create a map to store orders grouped by their warehouse
	ordersInWarehouse := make(map[int][]models.Order)
	// Create a map to store warehouse lat/long by WarehouseID
	warehouseLocation := make(map[int]struct {
		Lat  float64
		Long float64
	})

	// Group agents by their warehouse
	for _, agent := range agents {
		agentsInWarehouse[agent.WarehouseID] = append(agentsInWarehouse[agent.WarehouseID], agent)
	}

	// Group orders by their warehouse
	for _, order := range orders {
		ordersInWarehouse[order.WarehouseID] = append(ordersInWarehouse[order.WarehouseID], order)
	}

	// Store warehouse lat/long in the map
	for _, warehouse := range warehouses {
		warehouseLocation[warehouse.WarehouseID] = struct {
			Lat  float64
			Long float64
		}{
			Lat:  warehouse.Latitude,
			Long: warehouse.Longitude,
		}
	}

	// Initialize agent activity
	err := InitializeAgentActivity(agents)
	if err != nil {
		log.Println("Error initializing activity log:", err)
	}

	// Allocate orders to agents while considering conditions
	for warehouseID, orders := range ordersInWarehouse {
		agentsByWarehouse := agentsInWarehouse[warehouseID]

		// Get the current warehouse lat/long
		warehouseLatLong, ok := warehouseLocation[warehouseID]
		if !ok {
			log.Println("Warehouse location not found for warehouse:", warehouseID)
			continue
		}

		// Check if there are agents available in the warehouse
		if len(agentsByWarehouse) == 0 {
			log.Println("No agents found for warehouse:", warehouseID)
			continue
		}

		// Assign orders to agents
		ordersIndex := 0
		for _, order := range orders {
			lat := order.DestinationLat
			long := order.DestinationLong

			// Access warehouse lat/long
			fmt.Printf("Warehouse ID: %d, Lat: %f, Long: %f\n", warehouseID, warehouseLatLong.Lat, warehouseLatLong.Long)

			// Process order assignment logic here

			// Update index for next order
			ordersIndex++
		}
	}

	log.Println("Order allocation completed.")
	return dtos.Response{}
}
}


/*
* The flow of the allocation logic for assigning 60 orders to all agents:
*
* 1. Fetch all agents from the database.
* 2. Fetch all orders from the database
* 2. For each agent, identify the warehouse they are assigned to.
* 3. For each order, identify the warehouse they are assigned to.
* 3. Generate or retrieve 60 orders that need to be assigned to that agent.
* 4. For each order:
*   - Assign the agent's ID to the order.
*   - Assign the agent's warehouse ID to the order.
*   - Randomly generate delivery addresses or pull from a pool of available addresses.
*   - Set other necessary details like latitude/longitude of the delivery destination.
* 5. Save each order in the database with the respective agent's ID and warehouse ID.
* 6. Repeat for each agent until all orders are assigned.
 */
