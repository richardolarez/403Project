package models

type PharmacyTechnician struct {
	Name       string
	EmployeeID int
	password   string
	firstLogin bool
}

// Constructor
func NewPharmacyTechnician(name string, employeeID int) *PharmacyTechnician {
	return &PharmacyTechnician{
		Name:       name,
		EmployeeID: employeeID,
		password:   "defaultpassword",
		firstLogin: false,
	}
}

// func (p *PharmacyTechnician) Set...() {

// }
