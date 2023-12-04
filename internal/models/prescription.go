// models/inventory.go

package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// InventoryItem represents an item in the pharmacy's inventory.
type Prescription struct {
	ID           int     // Unique identifier for the item
	Drug         string  // Name of the item
	Doses        int     // Number of drug doses
	Strength     string  // Strength of the drug per dose
	Price        float64 // Price of the item
	Doctor       string  // Name of the doctor who prescribed the drug
	CustomerID   int     // ID of the customer who the prescription s assigned to
	IsFilled     bool    // Is the prescription filled?
	PharmacistID int     // ID of the pharmacist who filled the prescription
}

// NewInventoryItem creates a new InventoryItem object with the specified properties and a new ID.
func NewPrescription(id int, drug string, doses int, strength string, price float64, doctor string, customerid int) *Prescription {
	return &Prescription{
		ID:           id,
		Drug:         drug,
		Doses:        doses,
		Strength:     strength,
		Price:        price,
		Doctor:       doctor,
		CustomerID:   customerid,
		IsFilled:     false,
		PharmacistID: 0,
	}
}

// GetPrecriptions retrieves all prescription items.
func GetPrescriptions() ([]*Prescription, error) {
	// Read the contents of the database file
	data, err := ioutil.ReadFile("./db/database.json")
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into a map
	var db map[string]interface{}
	err = json.Unmarshal(data, &db)
	if err != nil {
		return nil, err
	}

	// Get the prescriptions object from the map'
	prescriptionsObj, ok := db["prescriptions"]
	if !ok {
		return nil, fmt.Errorf("employees object not found in database")
	}

	// Convert the prescriptions object to a JSON string
	prescriptionsJSON, err := json.Marshal(prescriptionsObj)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into an array of Prescription objects
	var prescriptions []*Prescription
	err = json.Unmarshal(prescriptionsJSON, &prescriptions)
	if err != nil {
		return nil, err
	}

	// Return the array of Prescription objects
	return prescriptions, nil
}

// AddPrescription adds a new prescription to the prescriptions db
func AddPrescription(id int, drug string, doses int, strength string, price float64, doctor string, customerid int) error {
	// Read the inventory data from the JSON file
	data, err := ioutil.ReadFile("./db/database.json")
	if err != nil {
		return fmt.Errorf("error reading prescriptions data: %v", err)
	}

	// Unmarshal the prescriptions data into a map
	var prescriptionsData map[string]interface{}
	err = json.Unmarshal(data, &prescriptionsData)
	if err != nil {
		return fmt.Errorf("error unmarshaling prescriptions data: %v", err)
	}

	// Get the prescriptions array from the data map
	prescriptionsArray, ok := prescriptionsData["prescriptions"].([]interface{})
	if !ok {
		return fmt.Errorf("error getting inventory array from data")
	}

	// Add the new item to the inventory array
	itemMap := map[string]interface{}{
		"id":           id,
		"drug":         drug,
		"doses":        doses,
		"strength":     strength,
		"price":        price,
		"doctor":       doctor,
		"customerID":   customerid,
		"isFilled":     false,
		"pharmacistID": 0,
	}

	prescriptionsArray = append(prescriptionsArray, itemMap)

	// Update the prescriptions data map
	prescriptionsData["prescriptions"] = prescriptionsArray

	// Marshal the prescription data back to JSON
	newData, err := json.Marshal(prescriptionsData)
	if err != nil {
		return fmt.Errorf("error marshaling prescriptions data: %v", err)
	}

	// Write the updated prescriptions data to the JSON file
	err = ioutil.WriteFile("database.json", newData, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error writing prescriptions data: %v", err)
	}

	return nil
}
