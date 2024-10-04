package services

import (
	"delivery-management-system/dtos"
	"delivery-management-system/initializers"
	"delivery-management-system/models"
	"log"
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

	// Create a map to store agents grouped by their warehouse
	agentsWarehouse := make(map[int][]models.Agent)
	// Create a map to store orders grouped by their warehouse
	ordersWarehouse := make(map[int][]models.Order)

	// Group agents by their warehouse
	for _, agent := range agents {
		agentsWarehouse[agent.WarehouseID] = append(agentsWarehouse[agent.WarehouseID], agent)
	}

	// Group orders by their warehouse
	for _, order := range orders {
		ordersWarehouse[order.WarehouseID] = append(ordersWarehouse[order.WarehouseID], order)
	}
	err := InitializeAgentActivity(agents)
	if err != nil {
		log.Println("Error initializing activity log:", err)
	}

	/*Conditions to allocate orders to agents-
	1. Agents cannot work for more than 10 hours in a day.
	2. Agents cannot drive more than 100 km in a day.
	*/
	//Allocate orders to agents in the same warehouse wile considering conditions above
	for warehouseID, orders := range ordersWarehouse {
		//omoo wetin be "for loop" again

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
