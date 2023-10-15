package dbinitializer

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/richardolarez/403Project/internal/models"
)

func InitializeDatabase() error {
	// Create sample Pharmacy instances using the constructor
	pharmacy1 := models.NewPharmacy("Pharmacy A", "Location A", true, "pharmacyA.com", "Owner A", 1234567890, "9:00 AM - 5:00 PM")
	pharmacy2 := models.NewPharmacy("Pharmacy B", "Location B", false, "pharmacyB.com", "Owner B", 9876543210, "10:00 AM - 6:00 PM")

	// Create an array of Pharmacy instances
	pharmacies := []*models.Pharmacy{pharmacy1, pharmacy2}

	// Create an array of InventoryItems
	items := []*models.InventoryItem{
		{
			ID:          1,
			Name:        "Item 1",
			Description: "Description for Item 1",
			Price:       10.99,
			Quantity:    5,
		},
		{
			ID:          2,
			Name:        "Item 2",
			Description: "Description for Item 2",
			Price:       15.99,
			Quantity:    10,
		},
		{
			ID:          3,
			Name:        "Medicine 1",
			Description: "Description for Medicine 1",
			Price:       20.99,
			Quantity:    15,
		},
		{
			ID:          4,
			Name:        "Medicine 2",
			Description: "Description for Medicine 2",
			Price:       25.99,
			Quantity:    20,
		},
	}

	// Create sample employees and append a new employee to the employees slice
	richardo := models.NewEmployee("richardo", "password123", "Richardo", "Larez", "Cashier")
	evan := models.NewEmployee("evan", "password123", "Evan", "Martin", "Pharmacist")
	javier := models.NewEmployee("javier", "password123", "Javier", "Cota", "Manager")
	carlos := models.NewEmployee("carlos", "password123", "Carlos", "Robles", "Pharmacist")
	mario := models.NewEmployee("mario", "password123", "Mario", "Weiler", "Cashier")
	antony := models.NewEmployee("antony", "password123", "Antony", "Mangala", "Cashier")

	// Create an array of employees
	employees := []*models.Employee{richardo, evan, javier, carlos, mario, antony}

	// Create a map to hold the data
	data := map[string]interface{}{
		"pharmacies": pharmacies,
		"items":      items,
		"employees":  employees,
	}

	// Marshal the data to JSON
	dataJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// Write the JSON data to the database file
	err = os.WriteFile("./db/database.json", dataJSON, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Database initialized successfully")
	return nil
}
