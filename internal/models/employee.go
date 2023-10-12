package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Employee represents an employee with basic information and a role.
type Employee struct {
	ID        int    // Unique identifier for the employee
	Username  string // Username for the employee's account
	Password  string // Password for the employee's account
	FirstName string // First name of the employee
	LastName  string // Last name of the employee
	Role      string // Role of the employee (e.g., Manager, Sales Associate, etc.)
}

// NewEmployee creates a new Employee object with the specified properties.
func NewEmployee(username, password, firstName, lastName, role string) *Employee {
	return &Employee{
		Username:  username,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
		Role:      role,
	}
}

// GetEmployeeByID retrieves an employee by ID.
func GetEmployeeByID(id int) (*Employee, error) {
	// Read the contents of the database file
	data, err := ioutil.ReadFile("database.json")
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into an array of Employee objects
	var employees []*Employee
	err = json.Unmarshal(data, &employees)
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

// GetAllEmployees retrieves all employees from the database.
func GetAllEmployees() ([]*Employee, error) {
	// Read the contents of the database file
	data, err := ioutil.ReadFile("database.json")
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into an array of Employee objects
	var employees []*Employee
	err = json.Unmarshal(data, &employees)
	if err != nil {
		return nil, err
	}

	return employees, nil
}
