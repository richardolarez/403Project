package models

type Pharmacist struct {
	name string
    employeeID int
}

// Constructor
func NewPharmacist(name, employeeID) *Pharmacist {
    return &Pharmacist{
        name:     name,
        employeeID: employeeID,
    }
}


// func (p *Pharmacist) Set...() {
    
// }
