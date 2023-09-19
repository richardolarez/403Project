package models

type Patient struct {
	name string
    patientID int
}

// Constructor
func NewPatient(name, patientID) *Patient {
    return &Patient{
        name:     name,
        patientID: patientID,
    }
}


// func (p *Patient) Set...() {
    
// }
