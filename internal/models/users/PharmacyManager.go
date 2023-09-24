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
