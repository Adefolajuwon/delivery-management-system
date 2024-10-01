package services

import (
	"delivery-management-system/initializers"
	models "delivery-management-system/models"
)

func allocate_orders(id string) {
	agent_warehouse := initializers.DB.First(&models.Agent{}, id)

}

/*
THE FLOW SHOULD BE SOMETHING LIKE THIS

1. Get Agent ID:

2. Retrieve Agent Information:
   - Query the database to get details of the agent using the provided Agent ID.
   - Check if the agent is active and eligible for allocation.

3. Get Warehouse Information:
   - Query the database to get the warehouse associated with the agent.
     - This could involve joining tables if necessary (e.g., Agents table and Warehouses table).
   - Retrieve the Warehouse ID, Name, and Location.
   - Validate if the warehouse is operational and has capacity for more agents.

4. Allocate order to Agents:
   - Perform the allocation logic:
     - Get all the unassigend orders in warehiuse
     -
     - Track the allocation date/time and any other relevant information.

5. Return Response:
   - Provide feedback to the user, confirming the allocation details.
   - Return the success status and any additional relevant information.
*/
