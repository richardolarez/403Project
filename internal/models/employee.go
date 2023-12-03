package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

// Employee represents an employee with basic information and a role.
type Employee struct {
	ID              int    // Unique identifier for the employee
	Username        string // Username for the employee's account
	Password        string // Password for the employee's account
	FirstName       string // First name of the employee
	LastName        string // Last name of the employee
	Role            string // Role of the employee (e.g., Manager, Sales Associate, etc.)
	RequiresNewPass bool
}

// NewEmployee creates a new Employee object with the specified properties and a new ID.
func NewEmployee(username, password, firstName, lastName, role string, requirenewpass bool) *Employee {
	rand.New(rand.NewSource(time.Now().UnixNano())) // Seed the random number generator
	id := rand.Intn(1000000)
	return &Employee{
		ID:              id,
		Username:        username,
		Password:        password,
		FirstName:       firstName,
		LastName:        lastName,
		Role:            role,
		RequiresNewPass: true,
	}
}

// GetEmployeeByID retrieves an employee by ID.
func GetEmployeeByID(id int) (*Employee, error) {
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

	// Get the employees object from the map
	employeesObj, ok := db["employees"]
	if !ok {
		return nil, fmt.Errorf("employees object not found in database")
	}

	// Convert the employees object to a JSON string
	employeesJSON, err := json.Marshal(employeesObj)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into an array of Employee objects
	var employees []*Employee
	err = json.Unmarshal(employeesJSON, &employees)
	if err != nil {
		return nil, err
	}

	// Find the employee with the specified ID
	for _, employee := range employees {
		if employee.ID == id {
			return employee, nil
		}
	}

	// If no employee was found, return an error
	return nil, fmt.Errorf("employee with ID %d not found", id)
}

// GetAllEmployees retrieves all employees.
func GetAllEmployees() ([]*Employee, error) {
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

	// Get the employees object from the map
	employeesObj, ok := db["employees"]
	if !ok {
		return nil, fmt.Errorf("employees object not found in database")
	}

	// Convert the employees object to a JSON string
	employeesJSON, err := json.Marshal(employeesObj)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into an array of Employee objects
	var employees []*Employee
	err = json.Unmarshal(employeesJSON, &employees)
	if err != nil {
		return nil, err
	}

	// Return the array of Employee objects
	return employees, nil
}

// DeleteEmployee deletes an employee by ID and first name from the database.
func DeleteEmployee(id int, FirstName string) error {
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

	// Get the employees object from the map
	employeesObj, ok := db["employees"]
	if !ok {
		return fmt.Errorf("employees object not found in database")
	}

	// Convert the employees object to a JSON string
	employeesJSON, err := json.Marshal(employeesObj)
	if err != nil {
		return err
	}

	// Unmarshal the JSON data into an array of Employee objects
	var employees []*Employee
	err = json.Unmarshal(employeesJSON, &employees)
	if err != nil {
		return err
	}

	// Find the index of the employee with the specified ID and first name
	index := -1
	for i, employee := range employees {
		if employee.ID == id && employee.FirstName == FirstName {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("employee with ID %d and first name %s not found", id, FirstName)
	}

	// Remove the employee from the array
	employees = append(employees[:index], employees[index+1:]...)

	// Update the employees object in the map
	db["employees"] = employees

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

// AddEmployee adds a new employee to the database.
func AddEmployee(username, password, firstName, lastName, role string) (*Employee, error) {
	// Generate a new unique ID for the employee
	rand.New(rand.NewSource(time.Now().UnixNano()))
	id := rand.Intn(1000000)

	// Create the new employee
	newEmployee := &Employee{
		ID:              id,
		Username:        username,
		Password:        password,
		FirstName:       firstName,
		LastName:        lastName,
		Role:            role,
		RequiresNewPass: true,
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

	// Get the employees object from the map
	employeesObj, ok := db["employees"]
	if !ok {
		return nil, fmt.Errorf("employees object not found in database")
	}

	// Convert the employees object to a JSON string
	employeesJSON, err := json.Marshal(employeesObj)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into an array of Employee objects
	var employees []*Employee
	err = json.Unmarshal(employeesJSON, &employees)
	if err != nil {
		return nil, err
	}

	// Append the new employee to the array
	employees = append(employees, newEmployee)

	// Update the employees object in the map
	db["employees"] = employees

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

	return newEmployee, nil
}

func UpdatePassword(username, oldPassword, newPassword string) error {
	// Read the employees data from the JSON file
	data, err := ioutil.ReadFile("./db/database.json")
	if err != nil {
		return err
	}

	// Unmarshal the employees data into a map
	var db map[string]interface{}
	err = json.Unmarshal(data, &db)
	if err != nil {
		return err
	}

	/// Get the employees object from the map
	employeesObj, ok := db["employees"]
	if !ok {
		return fmt.Errorf("employees object not found in database")
	}

	// Convert the employees object to a JSON string
	employeesJSON, err := json.Marshal(employeesObj)
	if err != nil {
		return err
	}

	// Unmarshal the JSON data into an array of Employee objects
	var employees []*Employee
	err = json.Unmarshal(employeesJSON, &employees)
	if err != nil {
		return err
	}

	// Find the employee with the provided username
	for _, employee := range employees {
		if employee.Username == username {
			// Check if the old password matches the employee's current password
			if employee.Password == oldPassword {
				// Update the password to the new password
				employee.Password = newPassword
				employee.RequiresNewPass = false

				// Update the employees array in the data map
				db["employees"] = employees

				// Marshal the updated data back to JSON
				updatedData, err := json.Marshal(db)
				if err != nil {
					return fmt.Errorf("error marshaling updated employees data: %v", err)
				}

				// Write the updated data back to the JSON file
				err = ioutil.WriteFile("./db/database.json", updatedData, 0644)
				if err != nil {
					return fmt.Errorf("error writing updated employees data: %v", err)
				}

				return nil
			} else {
				return fmt.Errorf("invalid old password")
			}
		}
	}

	// If no employee with the provided username is found, return an error
	return fmt.Errorf("invalid username")
}
