package models

import (
    "testing"
)

func TestNewPharmacy(t *testing.T) {
    name := "Test Pharmacy"
    location := "Test Location"
    isOpen := true

    pharmacy := NewPharmacy(name, location, isOpen)

    if pharmacy.name != name {
        t.Errorf("Expected name to be %s, but got %s", name, pharmacy.name)
    }

    if pharmacy.location != location {
        t.Errorf("Expected location to be %s, but got %s", location, pharmacy.location)
    }

    if pharmacy.isOpen != isOpen {
        t.Errorf("Expected isOpen to be %v, but got %v", isOpen, pharmacy.isOpen)
    }
}

func TestPharmacySetIsOpen(t *testing.T) {
    pharmacy := NewPharmacy("Test Pharmacy", "Test Location", false)

    pharmacy.SetIsOpen(true)

    if !pharmacy.isOpen {
        t.Errorf("Expected isOpen to be true, but got false")
    }
}

func TestPharmacyGetIsOpen(t *testing.T) {
    pharmacy := NewPharmacy("Test Pharmacy", "Test Location", true)

    isOpen := pharmacy.GetIsOpen()

    if !isOpen {
        t.Errorf("Expected GetIsOpen to return true, but got false")
    }
}
