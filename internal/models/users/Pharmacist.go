package models

type Pharmacist struct {
	Name       string
	EmployeeID int
	password   string
	firstLogin bool
}

// Constructor
func NewPharmacist(name string, employeeID int) *Pharmacist {
	return &Pharmacist{
		Name:       name,
		EmployeeID: employeeID,
		password:   "defaultpassword",
		firstLogin: false,
	}
}

// func (p *Pharmacist) Set...() {

// }
