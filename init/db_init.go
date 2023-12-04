package dbinitializer

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/richardolarez/403Project/internal/models"
)

func InitializeDatabase() error {
	// Create sample Pharmacy instances using the constructor
	pharmacy1 := models.NewPharmacy("Pharmacy A", "Location A", true, "pharmacyA.com", "Owner A", 1234567890, "9:00 AM - 5:00 PM")
	pharmacy2 := models.NewPharmacy("Pharmacy B", "Location B", false, "pharmacyB.com", "Owner B", 9876543210, "10:00 AM - 6:00 PM")

	// Create an array of Pharmacy instances
	pharmacies := []*models.Pharmacy{pharmacy1, pharmacy2}

	// Create an array of InventoryItems
	inventory := []*models.InventoryItem{
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
	richardo := models.NewEmployee("richardo", "password123", "Richardo", "Larez", "Manager", true)
	evan := models.NewEmployee("evan", "password123", "Evan", "Martin", "Pharmacist", true)
	javier := models.NewEmployee("javier", "password123", "Javier", "Cota", "Manager", true)
	carlos := models.NewEmployee("carlos", "password123", "Carlos", "Robles", "Pharmacist", true)
	mario := models.NewEmployee("mario", "password123", "Mario", "Weiler", "Cashier", true)
	antony := models.NewEmployee("antony", "password123", "Antony", "Mangala", "Pharmacist Technician", true)

	// Create an array of employees
	employees := []*models.Employee{richardo, evan, javier, carlos, mario, antony}

	//Create sameple customers and append new customer to the customers slice
	Richardo := models.NewCustomer("Richardo", "Larez", "01/01/2000", "somebody@gmail.com", "1234567890", "1234 Main St", "Blue Cross Blue Shield")
	Evan := models.NewCustomer("Evan", "Martin", "01/01/2000", "somebody@gmail.com", "1234567890", "1234 Main St", "Blue Cross Blue Shield")
	Javier := models.NewCustomer("Javier", "Cota", "01/01/2000", "somebody@gmail.com", "1234567890", "1234 Main St", "Blue Cross Blue Shield")
	Carlos := models.NewCustomer("Carlos", "Robles", "01/01/2000", "somebody@gmail.com", "1234567890", "1234 Main St", "Blue Cross Blue Shield")
	Mario := models.NewCustomer("Mario", "Weiler", "01/01/2000", "somebody@gmail.com", "1234567890", "1234 Main St", "Blue Cross Blue Shield")
	Antony := models.NewCustomer("Antony", "Mangala", "01/01/2000", "somebody@gmail.com", "1234567890", "1234 Main St", "Blue Cross Blue Shield")

	customers := []*models.Customer{Richardo, Evan, Javier, Carlos, Mario, Antony}

	p1 := models.NewPrescription(10, "Adderall", 10, "1mg", 10.99, "Dr Kureka", 20, "No", 33333)
	prescriptions := []*models.Prescription{p1}

	// Create sample medicines
	medicine1 := models.Medicine{
		ID:             1,
		Drug:           "Adderall",
		Doses:          5,
		Strength:       "10mg",
		Price:          99.99,
		ExpirationDate: time.Now().AddDate(0, 0, 30), // Set expiration date to 30 days from now
	}

	medicine2 := models.Medicine{
		ID:             2,
		Drug:           "Claritin-D",
		Doses:          10,
		Strength:       "20mg",
		Price:          14.99,
		ExpirationDate: time.Now().AddDate(1, 0, 0), // Set expiration date to 2 years from now
	}

	// Create an array of medicines
	medicines := []*models.Medicine{&medicine1, &medicine2}

	// Create a map to hold the data
	data := map[string]interface{}{
		"pharmacies":    pharmacies,
		"inventory":     inventory,
		"prescriptions": prescriptions,
		"employees":     employees,
		"customers":     customers,
		"medicines":     medicines,
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
