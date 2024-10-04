package main

import (
	"delivery-management-system/controller"
	"delivery-management-system/initializers"

	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	initializers.ConnectDb()

}
func main() {
	X := gin.Default()

	X.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})
	X.GET("/alloc", controllers.AllocateOrders)

	X.Run(":8080")94-
}
