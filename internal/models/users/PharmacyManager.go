package models

type PharmacyManager struct {
	Name       string
	EmployeeID int
	password   string
	firstLogin bool
}

// Constructor
func NewPharmacyManager(name string, employeeID int) *PharmacyManager {
	return &PharmacyManager{
		Name:       name,
		EmployeeID: employeeID,
		password:   "defaultpassword",
		firstLogin: false,
	}
}

// func (p *PharmacyManager) Set...() {

// }

func (p *PharmacyManager) SetfirstLogin(firstLogin bool) {
	p.firstLogin = firstLogin
}

// Getter method to retrieve the firstLogin status of the PharmacyManager
func (p *PharmacyManager) GetfirstLogin() bool {
	return p.firstLogin
}
