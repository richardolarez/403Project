// accountmanager/accountmanager.go

package accountmanager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/SFWE403/UArizonaPharmacy/internal/models"
)

// AuthenticateEmployee authenticates an employee login with the provided username and password.
func AuthenticateEmployee(username, password string) (*models.Employee, error) {
	// Read the employees data from the JSON file
	employeesData, err := ioutil.ReadFile("employees.json")
	if err != nil {
		return nil, fmt.Errorf("error reading employees data: %v", err)
	}

	// Unmarshal the employees data into an array of Employee objects
	var employees []*models.Employee
	err = json.Unmarshal(employeesData, &employees)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling employees data: %v", err)
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
