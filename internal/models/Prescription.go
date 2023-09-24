package models

import (
	"models" // <- needs fixing, probably an init thing
	"time"
)

// Prescription represents a prescription for medication.
type Prescription struct {
	PrescriptionNumber int
	medicine           string
	dosage             int
	prescriptor        string
	datePrescribed     time.Time
	filled             bool
	pharmacist         models.Pharmacist
	dateFilled         time.Time
}
