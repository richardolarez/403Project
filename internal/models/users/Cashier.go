package models

type Cashier struct {
	name string
    employeeID int
}

// Constructor
func NewCashier(name, employeeID) *Cashier {
    return &Cashier{
        name:     name,
        employeeID: employeeID,
    }
}


// func (c *Cashier) Set...() {
    
// }
