package models

type Cashier struct {
	Name       string
	EmployeeID int
	password   string
	firstLogin bool
}

// Constructor
func NewCashier(name string, employeeID int) *Cashier {
	return &Cashier{
		Name:       name,
		EmployeeID: employeeID,
		password:   "defaultpassword",
		firstLogin: false,
	}
}

// func (c *Cashier) Set...() {

// }
