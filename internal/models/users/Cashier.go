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
func (p *Cashier) SetfirstLogin(firstLogin bool) {
	p.firstLogin = firstLogin
}

// Getter method to retrieve the firstLogin status of the Cashier
func (p *Cashier) GetfirstLogin() bool {
	return p.firstLogin
}
