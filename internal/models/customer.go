// customer/customer.go

package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

// Customer represents a customer with basic information and a transaction history.
type Customer struct {
	ID           int                 // Unique identifier for the customer
	FirstName    string              // First name of the customer
	LastName     string              // Last name of the customer
	DOB          string              // Date of birth of the customer
	Email        string              // Email address of the customer
	PhoneNumber  string              // Phone number of the customer
	Address      string              // Address of the customer
	Insurance    string              // Insurance plan of the customer
	Transactions []*SalesTransaction // List of sales transactions for the customer
  AssignedPrescriptions []*Prescription 
}

// NewCustomer creates a new Customer object with the specified properties and a new ID.
func NewCustomer(firstName, lastName, dob, email, phoneNumber, address, insurance string) *Customer {
	rand.New(rand.NewSource(time.Now().UnixNano())) // Seed the random number generator
	id := rand.Intn(1000000)
	return &Customer{
		ID:           id,
		FirstName:    firstName,
		LastName:     lastName,
		DOB:          dob,
		Email:        email,
		PhoneNumber:  phoneNumber,
		Address:      address,
		Insurance:    insurance,
		Transactions: []*SalesTransaction{},
	}
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

// GetAllCustomers retrieves all customers.
func GetAllCustomers() ([]*Customer, error) {
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

	// Get the customers object from the map
	customersObj, ok := db["customers"]
	if !ok {
		return nil, fmt.Errorf("customers object not found in database")
	}

	// Convert the customers object to a JSON string
	customersJSON, err := json.Marshal(customersObj)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into an array of Customer objects
	var customers []*Customer
	err = json.Unmarshal(customersJSON, &customers)
	if err != nil {
		return nil, err
	}

	// Return the array of Customer objects
	return customers, nil
}

// AddEmployee adds a new customer to the database.
func AddCustomer(firstName, lastName, dob, email, phoneNumber, address, insurance string) (*Customer, error) {
	// Generate a new unique ID for the employee
	rand.New(rand.NewSource(time.Now().UnixNano()))
	id := rand.Intn(1000000)

	// Create the new customer
	newCustomer := &Customer{
		ID:           id,
		FirstName:    firstName,
		LastName:     lastName,
		DOB:          dob,
		Email:        email,
		PhoneNumber:  phoneNumber,
		Address:      address,
		Insurance:    insurance,
		Transactions: []*SalesTransaction{},
	}

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

	// Get the customers object from the map
	customerObj, ok := db["customers"]
	if !ok {
		return nil, fmt.Errorf("customers object not found in database")
	}

	// Convert the customers object to a JSON string
	customerJSON, err := json.Marshal(customerObj)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into an array of Customer objects
	var customers []*Customer
	err = json.Unmarshal(customerJSON, &customers)
	if err != nil {
		return nil, err
	}

	// Append the new customer to the array
	customers = append(customers, newCustomer)

	// Update the customers object in the map
	db["customers"] = customers

	// Marshal the map back to JSON
	updatedData, err := json.Marshal(db)
	if err != nil {
		return nil, err
	}

	// Write the updated JSON data to the database file
	err = ioutil.WriteFile("./db/database.json", updatedData, 0644)
	if err != nil {
		return nil, err
	}

	return newCustomer, nil
}

// DeleteCustomer deletes a customer by ID and first name from the database.
func DeleteCustomer(id int, FirstName string) error {
	// Read the contents of the database file
	data, err := ioutil.ReadFile("./db/database.json")
	if err != nil {
		return err
	}

	// Unmarshal the JSON data into a map
	var db map[string]interface{}
	err = json.Unmarshal(data, &db)
	if err != nil {
		return err
	}

	// Get the customers object from the map
	customersObj, ok := db["customers"]
	if !ok {
		return fmt.Errorf("customers object not found in database")
	}

	// Convert the customers object to a JSON string
	customersJSON, err := json.Marshal(customersObj)
	if err != nil {
		return err
	}

	// Unmarshal the JSON data into an array of Customer objects
	var customers []*Customer
	err = json.Unmarshal(customersJSON, &customers)
	if err != nil {
		return err
	}

	// Find the index of the customer with the specified ID and first name
	index := -1
	for i, customer := range customers {
		if customer.ID == id && customer.FirstName == FirstName {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("customer with ID %d and first name %s not found", id, FirstName)
	}

	// Remove the customer from the array
	customers = append(customers[:index], customers[index+1:]...)

	// Update the customers object in the map
	db["customers"] = customers

	// Marshal the map back to JSON
	updatedData, err := json.Marshal(db)
	if err != nil {
		return err
	}

	// Write the updated JSON data to the database file
	err = ioutil.WriteFile("./db/database.json", updatedData, 0644)
	if err != nil {
		return err
	}

	return nil
}

// AddTransaction adds a sales transaction to the customer's transaction history.
func (c *Customer) AddTransaction(transaction *SalesTransaction) {
	c.Transactions = append(c.Transactions, transaction)
}

// AddPrescription adds a prescription to the customer's assigned prescriptions.
func (c *Customer) AddPrescription(prescription *Prescription) {
	//TODO
	c.AssignedPrescriptions = append(c.AssignedPrescriptions, prescription)
}
