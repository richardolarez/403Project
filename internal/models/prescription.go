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

// GetInventory retrieves all inventory items.
func GetPrescriptions() ([]*Prescription, error) {
	// Read the prescription data from the JSON file
	data, err := ioutil.ReadFile("./db/database.json")
	if err != nil {
		return nil, fmt.Errorf("error reading prescriptions data: %v", err)
	}

	// Unmarshal the prescriptions data into a map
	var prescriptionsData map[string]interface{}
	err = json.Unmarshal(data, &prescriptionsData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling inventory data: %v", err)
	}

	// Get the inventory array from the data map
	prescriptionsArray, ok := prescriptionsData["prescriptions"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("error getting prescriptions array from data")
	}

	// Convert the prescriptions array to an array of Prescription objects
	var prescriptions []*Prescription
	for _, itemData := range prescriptionsArray {
		itemJSON, err := json.Marshal(itemData)
		if err != nil {
			return nil, fmt.Errorf("error marshaling prescription data: %v", err)
		}
		var item Prescription
		err = json.Unmarshal(itemJSON, &item)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling prescription data: %v", err)
		}
		prescriptions = append(prescriptions, &item)
	}

	// Return the list of inventory items
	return prescriptions, nil
}

// NewPrescription adds a new prescription to the prescriptions db
func NewPrescription(id int, drug string, doses int, strength string, price float64, doctor string, customerid int) error {
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
		"customerid":   customerid,
		"isfilled":     false,
		"pharmacistid": 0,
	}
	prescriptionsArray = append(prescriptionsArray, itemMap)

	// Update the prescriptions data map
	prescriptionsData["prescription"] = prescriptionsArray

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
