// models/inventory.go
package models

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

// InventoryItem represents an item in the pharmacy's inventory.
type Medicine struct {
	ID             int       // Unique identifier for the item
	Drug           string    // Name of the item
	Doses          int       // Number of drug doses
	Strength       string    // Strength of the drug per dose
	Price          float64   // Price of the item
	ExpirationDate time.Time // Expiration date of the drug
}

// newMedicine creates a new instance of the Medicine struct.
func newMedicine(id int, drug string, doses int, strength string, price float64, expirationDate time.Time) *Medicine {
	return &Medicine{
		ID:             id,
		Drug:           drug,
		Doses:          doses,
		Strength:       strength,
		Price:          price,
		ExpirationDate: expirationDate,
	}
}

// GetMedicine retrieves all medicine items from the database.
func GetMedicine() ([]*Medicine, error) {
	// Read the contents of the database file
	data, err := ioutil.ReadFile("db/database.json")
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into a slice of Medicine structs
	var medicine []*Medicine
	err = json.Unmarshal(data, &medicine)
	if err != nil {
		return nil, err
	}

	return medicine, nil
}

// addMedicine adds a new Medicine to the database.
func addMedicine(medicine *Medicine) error {
	// Read the contents of the database file
	data, err := ioutil.ReadFile("db/database.json")
	if err != nil {
		return err
	}

	// Unmarshal the JSON data into a slice of Medicine structs
	var medicines []*Medicine
	err = json.Unmarshal(data, &medicines)
	if err != nil {
		return err
	}

	// Add the new medicine to the slice
	medicines = append(medicines, medicine)

	// Marshal the updated slice back into JSON data
	updatedData, err := json.Marshal(medicines)
	if err != nil {
		return err
	}

	// Write the updated JSON data back to the database file
	err = ioutil.WriteFile("db/database.json", updatedData, 0644)
	if err != nil {
		return err
	}

	return nil
}
