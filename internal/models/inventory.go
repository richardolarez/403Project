// models/inventory.go

package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
func NewInventoryItem(id int, name string, description string, price float64, quantity int) error {
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

	// Add the new item to the inventory array
	itemMap := map[string]interface{}{
		"id":          id,
		"name":        name,
		"description": description,
		"price":       price,
		"quantity":    quantity,
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
	err = ioutil.WriteFile("database.json", newData, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error writing inventory data: %v", err)
	}

	return nil
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

// TotalValue returns the total value of the item in stock (price * quantity).
func (item *InventoryItem) TotalValue() float64 {
	return item.Price * float64(item.Quantity)
}
