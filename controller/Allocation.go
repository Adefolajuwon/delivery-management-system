package controllers

import (
	"delivery-management-system/services"

	"github.com/gin-gonic/gin"
)

func AllocateOrders(c *gin.Context) {
	cat  := services.AllocateOrders() // Now this will work correctly

	// if err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error()})
	// 	return
	// }

	// Respond with the allocated orders
	c.JSON(200, gin.H{"allocated_orders": cat})
}
