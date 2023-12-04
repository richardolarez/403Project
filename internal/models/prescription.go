// models/inventory.go

package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	IsFilled     string  // Is the prescription filled?
	PharmacistID int     // ID of the pharmacist who filled the prescription
}

type Database struct {
	Prescriptions []*Prescription `json:"prescriptions"`
	Medicines     []*Medicine     `json:"medicines"`
}

// NewInventoryItem creates a new InventoryItem object with the specified properties and a new ID.
func NewPrescription(id int, drug string, doses int, strength string, price float64, doctor string, customerid int, isfilled string, pharmacistid int) *Prescription {
	return &Prescription{
		ID:           id,
		Drug:         drug,
		Doses:        doses,
		Strength:     strength,
		Price:        price,
		Doctor:       doctor,
		CustomerID:   customerid,
		IsFilled:     isfilled,
		PharmacistID: pharmacistid,
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
func AddPrescription(id int, drug string, doses int, strength string, price float64, doctor string, customerid int, isfilled string, pharmacistid int) error {
	// Read the inventory data from the JSON file
	data, err := ioutil.ReadFile("./db/database.json")
	if err != nil {
		return fmt.Errorf("error reading database data: %v", err)
	}

	// Unmarshal the database data into a Database struct
	var database Database
	err = json.Unmarshal(data, &database)
	if err != nil {
		return fmt.Errorf("error unmarshaling database data: %v", err)
	}

	// Find the medicine with the given ID and decrement its quantity
	for _, medicine := range database.Medicines {
		if medicine.ID == id {
			medicine.Quantity -= doses
			if medicine.Quantity < 0 {
				return fmt.Errorf("not enough stock for medicine ID %d", id)
			}
			break
		}
	}

	// // Add the new item to the inventory array
	// itemMap := map[string]interface{}{
	// 	"id":           id,
	// 	"drug":         drug,
	// 	"doses":        doses,
	// 	"strength":     strength,
	// 	"price":        price,
	// 	"doctor":       doctor,
	// 	"customerid":   customerid,
	// 	"isfilled":     isfilled,
	// 	"pharmacistid": pharmacistid,
	// }

	// Marshal the Database struct back into JSON
	data, err = json.Marshal(database)
	if err != nil {
		return fmt.Errorf("error marshaling database data: %v", err)
	}

	// Write the JSON data back to the database.json file
	err = ioutil.WriteFile("./db/database.json", data, 0644)
	if err != nil {
		return fmt.Errorf("error writing database data: %v", err)
	}

	return nil
}
