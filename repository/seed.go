package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"delivery-management-system/initializers"
	"delivery-management-system/models" // Adjust the import path to your models

	"gorm.io/gorm"
)

var agentNames = []string{
	"John", "Jane", "Michael", "Sarah", "David", "Emily", "Daniel", "Sophia",
	"Chris", "Jessica", "Tom", "Olivia", "Paul", "Isabella", "James", "Mia",
	"Robert", "Charlotte", "Joseph", "Amelia", "Juwon",
}

var addresses = []string{
	"123 Main St", "456 Oak Ave", "789 Maple St", "321 Pine St", "654 Cedar Ave",
	"987 Birch Blvd", "234 Elm St", "876 Walnut Dr", "543 Aspen Ln", "109 Beech St",
}

// Seed the database with initial data
func SeedDatabase(db *gorm.DB) {
	if db == nil {
		log.Fatal("Database connection is nil. Cannot seed the database.")
	}

	log.Println("Seeding database...")

	// Create a new random source with the current time as the seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Seed 10 warehouses
	for i := 1; i <= 10; i++ {
		warehouse := models.Warehouse{
			Name:     fmt.Sprintf("Warehouse %d", i),
			Location: fmt.Sprintf("City Area %d", i),
		}
		if err := db.Create(&warehouse).Error; err != nil {
			log.Printf("Error creating warehouse: %v\n", err)
			continue
		}

		// Seed 20 agents per warehouse
		for j := 1; j <= 20; j++ {
			agent := models.Agent{
				Name:        agentNames[r.Intn(len(agentNames))],                            // Use the local random generator
				Phone:       fmt.Sprintf("+91%d", r.Intn(9999999999-1000000000)+1000000000), // Random phone number
				WarehouseID: warehouse.WarehouseID,                                          // Reference the warehouse's auto-incrementing ID
			}
			if err := db.Create(&agent).Error; err != nil {
				log.Printf("Error creating agent: %v\n", err)
				continue
			}

			// Seed 60 orders per agent
			for k := 1; k <= 60; k++ {
				order := models.Order{
					DeliveryAddress: addresses[r.Intn(len(addresses))], // Use the local random generator
					DestinationLat:  getRandomLatitude(),               // Assuming this function generates random latitudes
					DestinationLong: getRandomLongitude(),              // Assuming this function generates random longitudes
					AssignedAgentID: agent.AgentID,                     // Reference the agent's auto-incrementing ID
					WarehouseID:     warehouse.WarehouseID,             // Reference the warehouse's auto-incrementing ID
				}
				if err := db.Create(&order).Error; err != nil {
					log.Printf("Error creating order: %v\n", err)
					continue
				}
			}
		}
	}
	log.Println("Database seeding completed.")
}

// Utility functions for random geo-coordinates (Latitude, Longitude)
func getRandomLatitude() float64 {
	return rand.Float64()*(23.6345-22.7196) + 22.7196 // Example range for a test city
}

func getRandomLongitude() float64 {
	return rand.Float64()*(88.3639-87.5906) + 87.5906 // Example range for a test city
}

func init() {
	initializers.ConnectDb() // Ensure this function sets up initializers.DB
}

func main() {
	// Call the SeedDatabase function
	SeedDatabase(initializers.DB)
}
