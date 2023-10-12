package dbinitializer

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/SFWE403/UArizonaPharmacy/internal/models"
)

// InitializeDatabase initializes a JSON database with sample data
func InitializeDatabase() error {
	// Create sample Pharmacy instances using the constructor
	pharmacy1 := models.NewPharmacy("Pharmacy A", "Location A", true, "pharmacyA.com", "Owner A", 1234567890, "9:00 AM - 5:00 PM")
	pharmacy2 := models.NewPharmacy("Pharmacy B", "Location B", false, "pharmacyB.com", "Owner B", 9876543210, "10:00 AM - 6:00 PM")

	// Create sample InventoryItems
	item1 := models.NewInventoryItem(1, "Item 1", "Description of Item 1", 10.99, 100)
	item2 := models.NewInventoryItem(2, "Item 2", "Description of Item 2", 5.99, 50)

	// Create sample Medicines
	medicine1 := models.NewInventoryItem(3, "Medicine 1", "Description of Medicine 1", 15.99, 200)
	medicine2 := models.NewInventoryItem(4, "Medicine 2", "Description of Medicine 2", 8.99, 75)

	// Create sample employees
	cashier := models.NewEmployee("johndoe", "password123", "John", "Doe", "Cashier")
	manager := models.NewEmployee("janedoe", "password123", "Jane", "Doe", "Manager")
	pharmacist := models.NewEmployee("bobsmith", "password123", "Bob", "Smith", "Pharmacist")

	// Create an array of Pharmacy instances
	pharmacies := []*models.Pharmacy{pharmacy1, pharmacy2}

	// Create an array of InventoryItems
	items := []*models.InventoryItem{item1, item2, medicine1, medicine2}

	// Create an array of employees
	employees := []*models.Employee{cashier, manager, pharmacist}

	// Serialize the pharmacies array to JSON
	pharmaciesData, err := json.MarshalIndent(pharmacies, "", "  ")
	if err != nil {
		return err
	}

	// Serialize the items array to JSON
	itemsData, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return err
	}

	// Serialize the employees array to JSON
	employeesData, err := json.MarshalIndent(employees, "", "  ")
	if err != nil {
		return err
	}

	// Write JSON data to respective files
	err = os.WriteFile("./db/pharmacies.json", pharmaciesData, 0644)
	if err != nil {
		return err
	}

	err = os.WriteFile("./db/items.json", itemsData, 0644)
	if err != nil {
		return err
	}

	err = os.WriteFile("./db/employees.json", employeesData, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Database initialized successfully")

	return nil
}
