package dbinitializer

import (
	"encoding/json"
	"fmt"
	"os"

	// Import the Pharmacy struct from pharmacy.go in the same directory
	"github.com/SFWE403/UArizonaPharmacy/internal/models"
)

// InitializeDatabase initializes a JSON database with sample data
func InitializeDatabase() error {
	// Create sample Pharmacy instances using the constructor
	pharmacy1 := models.NewPharmacy("Pharmacy A", "Location A", true, "pharmacyA.com", "Owner A", 1234567890, "9:00 AM - 5:00 PM")
	pharmacy2 := models.NewPharmacy("Pharmacy B", "Location B", false, "pharmacyB.com", "Owner B", 9876543210, "10:00 AM - 6:00 PM")

	// Create an array of Pharmacy instances
	pharmacies := []*models.Pharmacy{pharmacy1, pharmacy2}

	// Serialize the pharmacies array to JSON
	jsonData, err := json.MarshalIndent(pharmacies, "", "  ")
	if err != nil {
		return err
	}

	// Write JSON data to a file
	file, err := os.Create("./db/pharmacies.json")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	fmt.Println("JSON database initialized and saved to 'pharmacies.json'")
	return nil
}
