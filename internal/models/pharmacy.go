package models

type Pharmacy struct {
	name         string
	location     string
	isOpen       bool
	website      string
	owner        string
	phoneNumber  int
	workingHours string
	isClosed     bool
}

// Constructor function for creating a new Pharmacy instance
func NewPharmacy(name, location string, isOpen bool, website string, owner string, phoneNumber int, workingHours string, isClosed bool) *Pharmacy {
	return &Pharmacy{
		name:         name,
		location:     location,
		website:      website,
		owner:        owner,
		phoneNumber:  phoneNumber,
		workingHours: workingHours,
		isOpen:       isOpen,
		isClosed:     isClosed,
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

// Setter method to update the isOpen status of the Pharmacy
func (p *Pharmacy) SetIsClosed(isClosed bool) {
	p.isClosed = isClosed
}

// Getter method to retrieve the isOpen status of the Pharmacy
func (p *Pharmacy) GetIsClosed() bool {
	return p.isClosed
}
