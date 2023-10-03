package models

type Pharmacy struct {
	Name         string `json:"name"`
	Location     string `json:"location"`
	IsOpen       bool   `json:"isOpen"`
	Website      string `json:"website"`
	Owner        string `json:"owner"`
	PhoneNumber  int    `json:"phoneNumber"`
	WorkingHours string `json:"workingHours"`
}

// Constructor function for creating a new Pharmacy instance
func NewPharmacy(name, location string, isOpen bool, website string, owner string, phoneNumber int, workingHours string) *Pharmacy {
	return &Pharmacy{
		Name:         name,
		Location:     location,
		Website:      website,
		Owner:        owner,
		PhoneNumber:  phoneNumber,
		WorkingHours: workingHours,
		IsOpen:       isOpen,
	}
}

// Getter methods for all attributes

func (p *Pharmacy) GetName() string {
	return p.Name
}

func (p *Pharmacy) GetLocation() string {
	return p.Location
}

func (p *Pharmacy) GetWebsite() string {
	return p.Website
}

func (p *Pharmacy) GetOwner() string {
	return p.Owner
}

func (p *Pharmacy) GetPhoneNumber() int {
	return p.PhoneNumber
}

func (p *Pharmacy) GetWorkingHours() string {
	return p.WorkingHours
}

// Setter methods for all attributes

func (p *Pharmacy) SetName(name string) {
	p.Name = name
}

func (p *Pharmacy) SetLocation(location string) {
	p.Location = location
}

func (p *Pharmacy) SetWebsite(website string) {
	p.Website = website
}

func (p *Pharmacy) SetOwner(owner string) {
	p.Owner = owner
}

func (p *Pharmacy) SetPhoneNumber(phoneNumber int) {
	p.PhoneNumber = phoneNumber
}

func (p *Pharmacy) SetWorkingHours(workingHours string) {
	p.WorkingHours = workingHours
}

// Setter method to update the isOpen status of the Pharmacy
func (p *Pharmacy) SetIsOpen(isOpen bool) {
	p.IsOpen = isOpen
}

// Getter method to retrieve the isOpen status of the Pharmacy
func (p *Pharmacy) GetIsOpen() bool {
	return p.IsOpen
}
