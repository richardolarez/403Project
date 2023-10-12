package models

import "time"

// SalesTransaction represents a sales transaction with basic information and a list of items sold.
type SalesTransaction struct {
	TransactionID   int              // Unique identifier for the transaction
	CustomerID      int              // Name of the customer associated with the transaction
	TransactionDate time.Time        // Date and time of the transaction
	ItemsSold       []*InventoryItem // List of items sold in the transaction
	TotalAmount     float64          // Total amount of the transaction
	PaymentMethod   string           // Payment method used for the transaction
}
