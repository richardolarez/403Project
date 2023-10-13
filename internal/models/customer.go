// customer/customer.go

package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Customer represents a customer with basic information and a transaction history.
type Customer struct {
	ID           int                 // Unique identifier for the customer
	FirstName    string              // First name of the customer
	LastName     string              // Last name of the customer
	Email        string              // Email address of the customer
	PhoneNumber  string              // Phone number of the customer
	Address      string              // Address of the customer
	Transactions []*SalesTransaction // List of sales transactions for the customer
}

// GetCustomer retrieves a customer by ID.
func GetCustomer(id int) (*Customer, error) {
	// Read the customers data from the JSON file
	customersData, err := ioutil.ReadFile("./db/database.json")
	if err != nil {
		return nil, fmt.Errorf("error reading customers data: %v", err)
	}

	// Unmarshal the customers data into a map
	var data map[string]interface{}
	err = json.Unmarshal(customersData, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling customers data: %v", err)
	}

	// Get the customers array from the data map
	customersArray, ok := data["customers"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("error getting customers array from data")
	}

	// Convert the customers array to an array of Customer objects
	var customers []*Customer
	for _, customerData := range customersArray {
		customerJSON, err := json.Marshal(customerData)
		if err != nil {
			return nil, fmt.Errorf("error marshaling customer data: %v", err)
		}
		var customer Customer
		err = json.Unmarshal(customerJSON, &customer)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling customer data: %v", err)
		}
		customers = append(customers, &customer)
	}

	// Find the customer with the provided ID
	for _, customer := range customers {
		if customer.ID == id {
			return customer, nil
		}
	}

	// If no customer with the provided ID is found, return an error
	return nil, fmt.Errorf("customer not found")
}

// AddTransaction adds a sales transaction to the customer's transaction history.
func (c *Customer) AddTransaction(transaction *SalesTransaction) {
	c.Transactions = append(c.Transactions, transaction)
}
