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

func (p *Pharmacist) SetfirstLogin(firstLogin bool) {
	p.firstLogin = firstLogin
}

// Getter method to retrieve the firstLogin status of the Pharmacist
func (p *Pharmacist) GetfirstLogin() bool {
	return p.firstLogin
}
