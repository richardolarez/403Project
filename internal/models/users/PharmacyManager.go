package models

type PharmacyManager struct {
	name string
    employeeID int
}

// Constructor
func NewPharmacyManager(name, employeeID) *PharmacyManager {
    return &PharmacyManager{
        name:     name,
        employeeID: employeeID,
    }
}


// func (p *PharmacyManager) Set...() {
    
// }
