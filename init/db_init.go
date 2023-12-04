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
			Name:        "Toothpaste",
			Description: "Toothpaste for oral hygiene",
			Price:       2.99,
			Quantity:    50,
		},
		{
			ID:          2,
			Name:        "Shampoo",
			Description: "Shampoo for hair care",
			Price:       5.99,
			Quantity:    30,
		},
		{
			ID:          3,
			Name:        "Soap",
			Description: "Soap for body cleansing",
			Price:       1.99,
			Quantity:    40,
		},
		{
			ID:          4,
			Name:        "Toilet Paper",
			Description: "Toilet paper for personal hygiene",
			Price:       3.99,
			Quantity:    60,
		},
		{
			ID:          5,
			Name:        "Laundry Detergent",
			Description: "Detergent for laundry",
			Price:       8.99,
			Quantity:    25,
		},
		{
			ID:          6,
			Name:        "Hand Sanitizer",
			Description: "Hand sanitizer for hygiene",
			Price:       4.99,
			Quantity:    20,
		},
		{
			ID:          7,
			Name:        "Face Mask",
			Description: "Face mask for protection",
			Price:       1.99,
			Quantity:    100,
		},
		{
			ID:          8,
			Name:        "Band-Aids",
			Description: "Band-Aids for minor injuries",
			Price:       3.49,
			Quantity:    50,
		},
		{
			ID:          9,
			Name:        "Cough Syrup",
			Description: "Cough syrup for relief",
			Price:       6.99,
			Quantity:    30,
		},
		{
			ID:          10,
			Name:        "Vitamins",
			Description: "Vitamins for health",
			Price:       9.99,
			Quantity:    40,
		},
		{
			ID:          11,
			Name:        "Pain Reliever",
			Description: "Pain reliever for headaches",
			Price:       4.49,
			Quantity:    50,
		},
		{
			ID:          12,
			Name:        "Allergy Medicine",
			Description: "Medicine for allergies",
			Price:       7.99,
			Quantity:    20,
		},
		{
			ID:          13,
			Name:        "Antacid",
			Description: "Antacid for heartburn",
			Price:       3.99,
			Quantity:    30,
		},
		{
			ID:          14,
			Name:        "Sunscreen",
			Description: "Sunscreen for sun protection",
			Price:       10.99,
			Quantity:    15,
		},
		{
			ID:          15,
			Name:        "Lip Balm",
			Description: "Lip balm for dry lips",
			Price:       2.49,
			Quantity:    40,
		},
		{
			ID:          16,
			Name:        "Bread",
			Description: "Freshly baked bread",
			Price:       1.49,
			Quantity:    75,
		},
		{
			ID:          17,
			Name:        "Milk",
			Description: "Whole milk for daily consumption",
			Price:       2.29,
			Quantity:    50,
		},
		{
			ID:          18,
			Name:        "Eggs",
			Description: "Farm-fresh eggs",
			Price:       3.99,
			Quantity:    60,
		},
		{
			ID:          19,
			Name:        "Cheese",
			Description: "Cheddar cheese for snacking",
			Price:       4.49,
			Quantity:    35,
		},
		{
			ID:          20,
			Name:        "Chicken Breast",
			Description: "Boneless chicken breast",
			Price:       7.99,
			Quantity:    25,
		},
		{
			ID:          21,
			Name:        "Pasta",
			Description: "Durum wheat pasta",
			Price:       1.99,
			Quantity:    45,
		},
		{
			ID:          22,
			Name:        "Tomatoes",
			Description: "Fresh tomatoes for cooking",
			Price:       0.99,
			Quantity:    30,
		},
		{
			ID:          23,
			Name:        "Potatoes",
			Description: "Russet potatoes for various dishes",
			Price:       2.49,
			Quantity:    40,
		},
		{
			ID:          24,
			Name:        "Canned Beans",
			Description: "Canned beans for convenience",
			Price:       1.79,
			Quantity:    50,
		},
		{
			ID:          25,
			Name:        "Frozen Pizza",
			Description: "Ready-to-bake frozen pizza",
			Price:       5.99,
			Quantity:    20,
		},
		{
			ID:          26,
			Name:        "Fresh Fruits",
			Description: "Assorted fresh fruits",
			Price:       6.99,
			Quantity:    30,
		},
		{
			ID:          27,
			Name:        "Yogurt",
			Description: "Greek yogurt for a healthy snack",
			Price:       3.49,
			Quantity:    40,
		},
		{
			ID:          28,
			Name:        "Salmon Fillet",
			Description: "Fresh salmon fillet",
			Price:       9.99,
			Quantity:    15,
		},
		{
			ID:          29,
			Name:        "Cereal",
			Description: "Whole grain cereal for breakfast",
			Price:       4.99,
			Quantity:    35,
		},
		{
			ID:          30,
			Name:        "Olive Oil",
			Description: "Extra virgin olive oil",
			Price:       8.49,
			Quantity:    25,
		},
		{
			ID:          31,
			Name:        "Honey",
			Description: "Pure honey for sweetening",
			Price:       3.29,
			Quantity:    40,
		},
		{
			ID:          32,
			Name:        "Green Tea",
			Description: "Organic green tea leaves",
			Price:       2.99,
			Quantity:    50,
		},
		{
			ID:          33,
			Name:        "Ground Coffee",
			Description: "Medium roast ground coffee",
			Price:       6.49,
			Quantity:    30,
		},
		{
			ID:          34,
			Name:        "Peanut Butter",
			Description: "Smooth peanut butter",
			Price:       2.79,
			Quantity:    45,
		},
		{
			ID:          35,
			Name:        "Granola Bars",
			Description: "Nutritious granola bars",
			Price:       3.99,
			Quantity:    50,
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
		Quantity:       10,
		ExpirationDate: time.Now().AddDate(0, 0, 30), // Set expiration date to 30 days from now
	}

	medicine2 := models.Medicine{
		ID:             2,
		Drug:           "Claritin-D",
		Doses:          10,
		Strength:       "20mg",
		Price:          14.99,
		Quantity:       20,
		ExpirationDate: time.Now().AddDate(1, 0, 0), // Set expiration date to 2 years from now
	}

	// Sample Medicines
	medicine3 := models.Medicine{
		ID:             3,
		Drug:           "Ibuprofen",
		Doses:          30,
		Strength:       "200mg",
		Price:          5.49,
		Quantity:       30,
		ExpirationDate: time.Now().AddDate(1, 6, 0), // Set expiration date to 1.5 years from now
	}

	medicine4 := models.Medicine{
		ID:             4,
		Drug:           "Amoxicillin",
		Doses:          20,
		Strength:       "500mg",
		Price:          8.99,
		Quantity:       15,
		ExpirationDate: time.Now().AddDate(2, 0, 0), // Set expiration date to 2 years from now
	}

	medicine5 := models.Medicine{
		ID:             5,
		Drug:           "Lisinopril",
		Doses:          30,
		Strength:       "10mg",
		Price:          12.79,
		Quantity:       25,
		ExpirationDate: time.Now().AddDate(0, 11, 0), // Set expiration date to 11 months from now
	}

	medicine6 := models.Medicine{
		ID:             6,
		Drug:           "Omeprazole",
		Doses:          14,
		Strength:       "20mg",
		Price:          6.49,
		Quantity:       10,
		ExpirationDate: time.Now().AddDate(1, 3, 0), // Set expiration date to 1.25 years from now
	}

	medicine7 := models.Medicine{
		ID:             7,
		Drug:           "Zithromax",
		Doses:          5,
		Strength:       "250mg",
		Price:          19.99,
		Quantity:       5,
		ExpirationDate: time.Now().AddDate(1, 6, 0), // Set expiration date to 1.5 years from now
	}

	medicine8 := models.Medicine{
		ID:             8,
		Drug:           "Prednisone",
		Doses:          10,
		Strength:       "5mg",
		Price:          7.99,
		Quantity:       12,
		ExpirationDate: time.Now().AddDate(0, 9, 0), // Set expiration date to 9 months from now
	}

	medicine9 := models.Medicine{
		ID:             9,
		Drug:           "Metformin",
		Doses:          30,
		Strength:       "500mg",
		Price:          10.29,
		Quantity:       18,
		ExpirationDate: time.Now().AddDate(1, 0, 0), // Set expiration date to 1 year from now
	}

	medicine10 := models.Medicine{
		ID:             10,
		Drug:           "Hydrochlorothiazide",
		Doses:          20,
		Strength:       "25mg",
		Price:          14.49,
		Quantity:       30,
		ExpirationDate: time.Now().AddDate(0, 8, 0), // Set expiration date to 8 months from now
	}

	medicine11 := models.Medicine{
		ID:             11,
		Drug:           "Atorvastatin",
		Doses:          30,
		Strength:       "40mg",
		Price:          9.99,
		Quantity:       22,
		ExpirationDate: time.Now().AddDate(1, 0, 0), // Set expiration date to 1 year from now
	}

	medicine12 := models.Medicine{
		ID:             12,
		Drug:           "Ciprofloxacin",
		Doses:          10,
		Strength:       "500mg",
		Price:          15.79,
		Quantity:       8,
		ExpirationDate: time.Now().AddDate(0, 10, 0), // Set expiration date to 10 months from now
	}

	medicine13 := models.Medicine{
		ID:             13,
		Drug:           "Albuterol",
		Doses:          15,
		Strength:       "90mcg",
		Price:          22.49,
		Quantity:       15,
		ExpirationDate: time.Now().AddDate(1, 2, 0), // Set expiration date to 1.17 years from now
	}

	medicine14 := models.Medicine{
		ID:             14,
		Drug:           "Sertraline",
		Doses:          30,
		Strength:       "50mg",
		Price:          11.99,
		Quantity:       20,
		ExpirationDate: time.Now().AddDate(0, 7, 0), // Set expiration date to 7 months from now
	}

	medicine15 := models.Medicine{
		ID:             15,
		Drug:           "Levothyroxine",
		Doses:          30,
		Strength:       "100mcg",
		Price:          8.49,
		Quantity:       10,
		ExpirationDate: time.Now().AddDate(0, 11, 0), // Set expiration date to 11 months from now
	}

	// Create an array of medicines
	medicines := []*models.Medicine{&medicine1, &medicine2, &medicine3, &medicine4, &medicine5, &medicine6, &medicine7, &medicine8, &medicine9, &medicine10, &medicine11, &medicine12, &medicine13, &medicine14, &medicine15}

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
