package test

import (
	"testing"

	"github.com/richardolarez/403Project/internal/models"
)

func TestPharmacy(t *testing.T) {
	// Create a Pharmacy instance for testing
	pharmacy := models.NewPharmacy("Test Pharmacy", "Test Location", true, "test.com", "Test Owner", 1234567890, "9:00 AM - 5:00 PM")

	// Test getters and setters for all attributes
	t.Run("Test getters and setters", func(t *testing.T) {
		// Test getters
		if pharmacy.GetName() != "Test Pharmacy" {
			t.Error("Expected name to be 'Test Pharmacy'")
		}

		if pharmacy.GetLocation() != "Test Location" {
			t.Error("Expected location to be 'Test Location'")
		}

		if pharmacy.GetWebsite() != "test.com" {
			t.Error("Expected website to be 'test.com'")
		}

		if pharmacy.GetOwner() != "Test Owner" {
			t.Error("Expected owner to be 'Test Owner'")
		}

		if pharmacy.GetPhoneNumber() != 1234567890 {
			t.Error("Expected phoneNumber to be 1234567890")
		}

		if pharmacy.GetWorkingHours() != "9:00 AM - 5:00 PM" {
			t.Error("Expected workingHours to be '9:00 AM - 5:00 PM'")
		}

		// Test setters
		pharmacy.SetName("Updated Pharmacy Name")
		if pharmacy.GetName() != "Updated Pharmacy Name" {
			t.Error("Expected updated name to be 'Updated Pharmacy Name'")
		}

		pharmacy.SetLocation("Updated Location")
		if pharmacy.GetLocation() != "Updated Location" {
			t.Error("Expected updated location to be 'Updated Location'")
		}

		pharmacy.SetWebsite("updated.com")
		if pharmacy.GetWebsite() != "updated.com" {
			t.Error("Expected updated website to be 'updated.com'")
		}

		pharmacy.SetOwner("Updated Owner")
		if pharmacy.GetOwner() != "Updated Owner" {
			t.Error("Expected updated owner to be 'Updated Owner'")
		}

		pharmacy.SetPhoneNumber(987654321)
		if pharmacy.GetPhoneNumber() != 987654321 {
			t.Error("Expected updated phoneNumber to be 987654321")
		}

		pharmacy.SetWorkingHours("10:00 AM - 6:00 PM")
		if pharmacy.GetWorkingHours() != "10:00 AM - 6:00 PM" {
			t.Error("Expected updated workingHours to be '10:00 AM - 6:00 PM'")
		}
	})

	// Test isOpen getter and setter
	t.Run("Test isOpen getter and setter", func(t *testing.T) {
		if pharmacy.GetIsOpen() != true {
			t.Error("Expected isOpen to be true")
		}

		pharmacy.SetIsOpen(false)
		if pharmacy.GetIsOpen() != false {
			t.Error("Expected updated isOpen to be false")
		}
	})
}
