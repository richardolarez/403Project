package models

type Pharmacy struct {
	name         string
	location     string
	isOpen       bool
	website      string
	owner        string
	phoneNumber  int
	workingHours string
}

// Constructor function for creating a new Pharmacy instance
func NewPharmacy(name, location string, isOpen bool, website string, owner string, phoneNumber int, workingHours string) *Pharmacy {
	return &Pharmacy{
		name:         name,
		location:     location,
		website:      website,
		owner:        owner,
		phoneNumber:  phoneNumber,
		workingHours: workingHours,
		isOpen:       isOpen,
	}
}

// Getter methods for all attributes

func (p *Pharmacy) GetName() string {
	return p.name
}

func (p *Pharmacy) GetLocation() string {
	return p.location
}

func (p *Pharmacy) GetWebsite() string {
	return p.website
}

func (p *Pharmacy) GetOwner() string {
	return p.owner
}

func (p *Pharmacy) GetPhoneNumber() int {
	return p.phoneNumber
}

func (p *Pharmacy) GetWorkingHours() string {
	return p.workingHours
}

// Setter methods for all attributes

func (p *Pharmacy) SetName(name string) {
	p.name = name
}

func (p *Pharmacy) SetLocation(location string) {
	p.location = location
}

func (p *Pharmacy) SetWebsite(website string) {
	p.website = website
}

func (p *Pharmacy) SetOwner(owner string) {
	p.owner = owner
}

func (p *Pharmacy) SetPhoneNumber(phoneNumber int) {
	p.phoneNumber = phoneNumber
}

func (p *Pharmacy) SetWorkingHours(workingHours string) {
	p.workingHours = workingHours
}

// Setter method to update the isOpen status of the Pharmacy
func (p *Pharmacy) SetIsOpen(isOpen bool) {
	p.isOpen = isOpen
}

// Getter method to retrieve the isOpen status of the Pharmacy
func (p *Pharmacy) GetIsOpen() bool {
	return p.isOpen
}
