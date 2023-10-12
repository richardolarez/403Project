// employee/employee.go

package models

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
	// Retrieve the employee from the database (you can implement this logic)
	// For now, we'll just return a dummy employee object
	employee := &Employee{
		ID:        id,
		Username:  "johndoe",
		FirstName: "John",
		LastName:  "Doe",
		Role:      "Manager",
	}

	return employee, nil
}
