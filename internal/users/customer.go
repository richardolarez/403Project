// models/customer.go

package users

import (
	"time"

	"github.com/SFWE403/UArizonaPharmacy/internal/models"
	"github.com/SFWE403/UArizonaPharmacy/internal/service"
)

// Customer represents a customer with basic information and a history of transactions.
type Customer struct {
	ID           int                         // Unique identifier for the customer
	FirstName    string                      // First name of the customer
	LastName     string                      // Last name of the customer
	Email        string                      // Email address of the customer
	PhoneNumber  string                      // Phone number of the customer
	Address      string                      // Address of the customer
	Transactions []*service.SalesTransaction // List of sales transactions associated with the customer
}

// NewCustomer creates a new Customer object with the provided information.
func NewCustomer(id int, firstName, lastName, email, phoneNumber, address string) *Customer {
	return &Customer{
		ID:           id,
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		PhoneNumber:  phoneNumber,
		Address:      address,
		Transactions: []*service.SalesTransaction{},
	}
}

// AddTransaction adds a sales transaction to the customer's history.
func (c *Customer) AddTransaction(transaction *service.SalesTransaction) {
	c.Transactions = append(c.Transactions, transaction)

}

// GetCustomer retrieves a customer by ID.
func GetCustomer(id int) (*Customer, error) {
	// Retrieve the customer from the database (you can implement this logic)
	// For now, we'll just return a dummy customer object
	customer := &Customer{
		ID:          id,
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "johndoe@example.com",
		PhoneNumber: "555-1234",
		Address:     "123 Main St",
		Transactions: []*service.SalesTransaction{
			{
				TransactionID:   1,
				CustomerName:    "John Doe",
				TransactionDate: time.Now(),
				ItemsSold: []*models.InventoryItem{
					{Name: "Aspirin", Price: 4.99},
					{Name: "Ibuprofen", Price: 6.99},
				},
				TotalAmount:   11.98,
				PaymentMethod: "Credit card",
			},
			{
				TransactionID:   2,
				CustomerName:    "John Doe",
				TransactionDate: time.Now(),
				ItemsSold: []*models.InventoryItem{
					{Name: "Tylenol", Price: 5.99},
				},
				TotalAmount:   5.99,
				PaymentMethod: "Cash",
			},
		},
	}

	return customer, nil
}
