// models/customer.go

package models

import "time"

type Transaction struct {
	TransactionID   int
	CustomerName    string
	TransactionDate time.Time
	ItemsSold       []*InventoryItem
	TotalAmount     float64
	PaymentMethod   string
}

// Customer represents a customer with basic information and a history of transactions.
type Customer struct {
	ID           int            // Unique identifier for the customer
	FirstName    string         // First name of the customer
	LastName     string         // Last name of the customer
	Email        string         // Email address of the customer
	PhoneNumber  string         // Phone number of the customer
	Address      string         // Address of the customer
	Transactions []*Transaction // List of sales transactions associated with the customer
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
		Transactions: []*Transaction{},
	}
}

// AddTransaction adds a sales transaction to the customer's history.
func (c *Customer) AddTransaction(transaction *Transaction) {
	c.Transactions = append(c.Transactions, transaction)

}
