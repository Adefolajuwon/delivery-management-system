package services

import (
	"delivery-management-system/dtos"
	"delivery-management-system/helper"
	"delivery-management-system/initializers"
	"delivery-management-system/models"
	"log"
	"sort"
)

// AllocateOrders function that assigns the best agent to each order based on proximity and availability.
func AllocateOrders() dtos.Response {
	var agents []models.Agent
	if err := initializers.DB.Find(&agents).Error; err != nil {
		log.Println("Error fetching agents:", err)
	}

	var orders []models.Order
	if err := initializers.DB.Where("assigned_agent_id IS NULL").Find(&orders).Error; err != nil {
		log.Println("Error fetching orders:", err)
	}

	var warehouses []models.Warehouse
	if err := initializers.DB.Find(&warehouses).Error; err != nil {
		log.Println("Error fetching warehouses:", err)
	}

	// Create a map to store agents grouped by their warehouse
	agentsInWarehouse := make(map[int][]models.Agent)
	// Create a map to store orders grouped by their warehouse
	ordersInWarehouse := make(map[int][]models.Order)
	// Create a map to store warehouse location data
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

	// Store latitude and longitude for each warehouse
	for _, warehouse := range warehouses {
		warehouseLocation[warehouse.WarehouseID] = struct {
			Lat  float64
			Long float64
		}{
			Lat:  warehouse.Latitude,
			Long: warehouse.Longitude,
		}
	}

	// Loop through each warehouse and assign orders to agents
	for warehouseID, orders := range ordersInWarehouse {
		agentsByWarehouse := agentsInWarehouse[warehouseID]

		// Check if there are agents available in the warehouse
		if len(agentsByWarehouse) == 0 {
			log.Println("No agents found for warehouse:", warehouseID)
			continue
		}

		// Get the warehouse's lat and long
		warehouseLatLong := warehouseLocation[warehouseID]

		// Step 1: Sort the orders based on Haversine distance
		sort.Slice(orders, func(i, j int) bool {
			lat1, long1 := orders[i].DestinationLat, orders[i].DestinationLong
			lat2, long2 := orders[j].DestinationLat, orders[j].DestinationLong

			// Calculate the distance of both orders from the warehouse
			dist1 := helper.Haversine(warehouseLatLong.Lat, warehouseLatLong.Long, lat1, long1)
			dist2 := helper.Haversine(warehouseLatLong.Lat, warehouseLatLong.Long, lat2, long2)

			// Sort in ascending order of distance (i.e., closer orders first)
			return dist1 < dist2
		})

		// Step 2: Assign orders to the best available agent
		for _, order := range orders {
			bestAgent := findBestAgent(agentsByWarehouse, warehouseLatLong, order)

			if bestAgent != nil {
				// Assign the order to the best agent
				order.AssignedAgentID = bestAgent.AgentID
				initializers.DB.Save(&order)

				// Update the agent's workload (distance and working hours)
				bestAgent.Workload += calculateOrderTimeAndDistance(order)
				initializers.DB.Save(&bestAgent)
			} else {
				log.Println("No suitable agent available for order:", order.OrderID)
			}
		}
	}

	log.Println("Order allocation completed.")
	return dtos.Response{}
}

// Function to find the best agent for an order
func findBestAgent(agents []models.Agent, warehouseLatLong struct{ Lat, Long float64 }, order models.Order) *models.Agent {
	// Sort agents based on proximity to the order's destination
	sort.Slice(agents, func(i, j int) bool {
		agent1 := agents[i]
		agent2 := agents[j]

		// Calculate the distance between agent's current location and the order destination
		dist1 := helper.Haversine(agent1.Latitude, agent1.Longitude, order.DestinationLat, order.DestinationLong)
		dist2 := helper.Haversine(agent2.Latitude, agent2.Longitude, order.DestinationLat, order.DestinationLong)

		return dist1 < dist2
	})

	// Iterate over the sorted agents and find one who meets the workload and distance criteria
	for _, agent := range agents {
		// Check if the agent can handle more work (e.g., max hours and distance)
		if agent.Workload < maxDailyWorkload && agent.DailyDistance+calculateOrderDistance(agent, order) <= maxDailyDistance {
			return &agent
		}
	}

	// Return nil if no suitable agent is found
	return nil
}

// Utility function to calculate the distance of an order
func calculateOrderDistance(agent models.Agent, order models.Order) float64 {
	return helper.Haversine(agent.Latitude, agent.Longitude, order.DestinationLat, order.DestinationLong)
}

// Utility function to calculate the time and distance of an order
func calculateOrderTimeAndDistance(order models.Order) float64 {
	// Assuming some logic to calculate the time taken for the order based on the distance
	return order.Distance * timePerKm // Hypothetical time calculation
}

/*
	log.Println("Order allocation completed.")
	return dtos.Response{}
}

// Function to find the best agent for an order
func findBestAgent(agents []models.Agent, warehouseLatLong struct{ Lat, Long float64 }, order models.Order) *models.Agent {
	// Sort agents based on proximity to the order's destination
	sort.Slice(agents, func(i, j int) bool {
		agent1 := agents[i]
		agent2 := agents[j]

		// Calculate the distance between agent's current location and the order destination
		dist1 := helper.Haversine(agent1.Latitude, agent1.Longitude, order.DestinationLat, order.DestinationLong)
		dist2 := helper.Haversine(agent2.Latitude, agent2.Longitude, order.DestinationLat, order.DestinationLong)

		return dist1 < dist2
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
