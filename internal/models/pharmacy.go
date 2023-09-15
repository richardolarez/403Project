package models

import "fmt"

type Pharmacy struct {
    name     string
    location string
    isOpen   bool
}

// Constructor function for creating a new Pharmacy instance
func NewPharmacy(name, location string, isOpen bool) *Pharmacy {
    return &Pharmacy{
        name:     name,
        location: location,
        isOpen:   isOpen,
    }
}

// Setter method to update the isOpen status of the Pharmacy
func (p *Pharmacy) SetIsOpen(isOpen bool) {
    p.isOpen = isOpen
}

// Getter method to retrieve the isOpen status of the Pharmacy
func (p *Pharmacy) GetIsOpen() bool {
    return p.isOpen
}
