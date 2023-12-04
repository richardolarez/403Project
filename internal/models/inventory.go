// models/inventory.go

package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

// InventoryItem represents an item in the pharmacy's inventory.
type InventoryItem struct {
	ID          int     // Unique identifier for the item
	Name        string  // Name of the item
	Description string  // Description of the item
	Price       float64 // Price of the item
	Quantity    int     // Quantity of the item in stock
}

// GetInventory retrieves all inventory items.
func GetInventory() ([]*InventoryItem, error) {
	// Read the inventory data from the JSON file
	data, err := ioutil.ReadFile("db/database.json")
	if err != nil {
		return nil, fmt.Errorf("error reading inventory data: %v", err)
	}

	// Unmarshal the inventory data into a map
	var inventoryData map[string]interface{}
	err = json.Unmarshal(data, &inventoryData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling inventory data: %v", err)
	}

	// Get the inventory array from the data map
	inventoryArray, ok := inventoryData["inventory"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("error getting inventory array from data")
	}

	// Convert the inventory array to an array of InventoryItem objects
	var inventory []*InventoryItem
	for _, itemData := range inventoryArray {
		itemJSON, err := json.Marshal(itemData)
		if err != nil {
			return nil, fmt.Errorf("error marshaling inventory item data: %v", err)
		}
		var item InventoryItem
		err = json.Unmarshal(itemJSON, &item)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling inventory item data: %v", err)
		}
		inventory = append(inventory, &item)
	}

	// Return the list of inventory items
	return inventory, nil
}

// NewInventoryItem adds a new inventory item to the database.
func NewInventoryItem(name string, description string, price float64, quantity int, isPrescription bool) error {
	// Read the inventory data from the JSON file
	data, err := ioutil.ReadFile("./db/database.json")
	if err != nil {
		return fmt.Errorf("error reading inventory data: %v", err)
	}

	// Unmarshal the inventory data into a map
	var inventoryData map[string]interface{}
	err = json.Unmarshal(data, &inventoryData)
	if err != nil {
		return fmt.Errorf("error unmarshaling inventory data: %v", err)
	}

	// Get the inventory array from the data map
	inventoryArray, ok := inventoryData["inventory"].([]interface{})
	if !ok {
		return fmt.Errorf("error getting inventory array from data")
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	id := rand.Intn(1000000)
	// Add the new item to the inventory array
	itemMap := map[string]interface{}{
		"ID":             id,
		"Name":           name,
		"Description":    description,
		"Price":          price,
		"Quantity":       quantity,
		"IsPrescription": isPrescription,
	}
	inventoryArray = append(inventoryArray, itemMap)

	// Update the inventory data map
	inventoryData["inventory"] = inventoryArray

	// Marshal the inventory data back to JSON
	newData, err := json.Marshal(inventoryData)
	if err != nil {
		return fmt.Errorf("error marshaling inventory data: %v", err)
	}

	// Write the updated inventory data to the JSON file
	err = ioutil.WriteFile("./db/database.json", newData, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error writing inventory data: %v", err)
	}

	return nil
}

// GetInventoryItem retrieves an inventory item by ID.
func GetInventoryItem(id int) (*InventoryItem, error) {
	// Read the inventory data from the JSON file
	data, err := ioutil.ReadFile("./db/database.json")
	if err != nil {
		return nil, fmt.Errorf("error reading inventory data: %v", err)
	}

	// Unmarshal the inventory data into a map
	var inventoryData map[string]interface{}
	err = json.Unmarshal(data, &inventoryData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling inventory data: %v", err)
	}

	// Get the inventory array from the data map
	inventoryArray, ok := inventoryData["inventory"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("error getting inventory array from data")
	}

	// Find the item with the given ID
	for _, itemData := range inventoryArray {
		itemJSON, err := json.Marshal(itemData)
		if err != nil {
			return nil, fmt.Errorf("error marshaling inventory item data: %v", err)
		}
		var item InventoryItem
		err = json.Unmarshal(itemJSON, &item)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling inventory item data: %v", err)
		}
		if item.ID == id {
			return &item, nil
		}
	}

	// Return an error if the item is not found
	return nil, fmt.Errorf("item not found")
}

// IncreaseQuantity increases the quantity of the item in stock.
func (item *InventoryItem) IncreaseQuantity(amount int) {
	item.Quantity += amount
}

// DecreaseQuantity decreases the quantity of the item in stock.
func (item *InventoryItem) DecreaseQuantity(amount int) {
	if item.Quantity >= amount {
		item.Quantity -= amount
	}
}

// Update current inventory item
func (item *InventoryItem) Update(name string, description string, price float64, quantity int, isPrescription bool) error {
	// Update the item's fields
	item.Name = name
	item.Description = description
	item.Price = price
	item.Quantity = quantity

	// Read the inventory data from the JSON file
	data, err := ioutil.ReadFile("./db/database.json")
	if err != nil {
		return fmt.Errorf("error reading inventory data: %v", err)
	}

	// Unmarshal the inventory data into a map
	var inventoryData map[string]interface{}
	err = json.Unmarshal(data, &inventoryData)
	if err != nil {
		return fmt.Errorf("error unmarshaling inventory data: %v", err)
	}

	// Get the inventory array from the data map
	inventoryArray, ok := inventoryData["inventory"].([]interface{})
	if !ok {
		return fmt.Errorf("error getting inventory array from data")
	}

	// Find and update the item in the inventory array
	for i, itemData := range inventoryArray {
		itemJSON, err := json.Marshal(itemData)
		if err != nil {
			return fmt.Errorf("error marshaling inventory item data: %v", err)
		}
		var existingItem InventoryItem
		err = json.Unmarshal(itemJSON, &existingItem)
		if err != nil {
			return fmt.Errorf("error unmarshaling inventory item data: %v", err)
		}
		if existingItem.ID == item.ID {
			// Update the item in the inventory array
			inventoryArray[i] = item
			break
		}
	}

	// Update the inventory data map
	inventoryData["inventory"] = inventoryArray

	// Marshal the inventory data back to JSON
	newData, err := json.Marshal(inventoryData)
	if err != nil {
		return fmt.Errorf("error marshaling inventory data: %v", err)
	}

	// Write the updated inventory data to the JSON file
	err = ioutil.WriteFile("./db/database.json", newData, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error writing inventory data: %v", err)
	}

	return nil
}

// TotalValue returns the total value of the item in stock (price * quantity).
func (item *InventoryItem) TotalValue() float64 {
	return item.Price * float64(item.Quantity)
}
