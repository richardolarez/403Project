// accountmanager/accountmanager.go

package accountmanager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/richardolarez/403Project/internal/models"
)

// AuthenticateEmployee authenticates an employee login with the provided username and password.
func AuthenticateEmployee(username, password string) (*models.Employee, error) {
	// Read the employees data from the JSON file
	data, err := ioutil.ReadFile("./db/database.json")
	if err != nil {
		return nil, fmt.Errorf("error reading employees data: %v", err)
	}

	// Unmarshal the employees data into a map
	var employeesData map[string]interface{}
	err = json.Unmarshal(data, &employeesData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling employees data: %v", err)
	}

	// Get the employees array from the data map
	employeesArray, ok := employeesData["employees"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("error getting employees array from data")
	}

	// Convert the employees array to an array of Employee objects
	var employees []*models.Employee
	for _, employeeData := range employeesArray {
		employeeJSON, err := json.Marshal(employeeData)
		if err != nil {
			return nil, fmt.Errorf("error marshaling employee data: %v", err)
		}
		var employee models.Employee
		err = json.Unmarshal(employeeJSON, &employee)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling employee data: %v", err)
		}
		employees = append(employees, &employee)
	}

	// Find the employee with the provided username
	for _, employee := range employees {
		if employee.Username == username {
			// Check if the provided password matches the employee's password
			if employee.Password == password {
				return employee, nil
			} else {
				return nil, fmt.Errorf("invalid password")
			}
		}
	}

	// If no employee with the provided username is found, return an error
	return nil, fmt.Errorf("invalid username")
}
