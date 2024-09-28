package main

import (
	"github.com/Adefolajuwon/delivery-management-system/initializers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {

}
func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	router.Run(":8080")
}
