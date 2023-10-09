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

	// Create an array of Pharmacy instances
	pharmacies := []*models.Pharmacy{pharmacy1, pharmacy2}

	// Create an array of InventoryItems
	items := []*models.InventoryItem{item1, item2, medicine1, medicine2}

	// Create objects with arrays
	data := map[string]interface{}{
		"pharmacies": pharmacies,
		"items":      items,
	}

	// Serialize the data to JSON
	dataJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// Write JSON data to a single file
	dataFile, err := os.Create("./db/database.json")
	if err != nil {
		return err
	}
	defer dataFile.Close()

	_, err = dataFile.Write(dataJSON)
	if err != nil {
		return err
	}

	fmt.Println("JSON database initialized and saved to 'database.json'")
	return nil
}
