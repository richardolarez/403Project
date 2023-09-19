package models

type PharmacyTechnician struct {
	name string
    employeeID int
}

// Constructor
func NewPharmacyTechnician(name, employeeID) *PharmacyTechnician {
    return &PharmacyTechnician{
        name:     name,
        employeeID: employeeID,
    }
}


// func (p *PharmacyTechnician) Set...() {
    
// }
