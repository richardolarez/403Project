package models

import (
	"time"
)

type Patient struct {
	name        string
	patientID   int
	dateOfBirth time.Time
	address     string
	phoneNumber string
	email       string
	insurance   [2]string // 0 = provider/plan, 1 = insurance number

}

// Constructor
func NewPatient(name string, patientID int) *Patient {
	return &Patient{
		name:      name,
		patientID: patientID,
	}
}

// func (p *Patient) Set...() {

// }
