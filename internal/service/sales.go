// service/sales.go

package service

import (
	"time"

	"github.com/SFWE403/UArizonaPharmacy/internal/models"
)

// SalesTransaction represents a sales transaction.
type SalesTransaction struct {
	TransactionID   int
	CustomerName    string
	TransactionDate time.Time
	ItemsSold       []*models.InventoryItem
	TotalAmount     float64
	PaymentMethod   string
}

// Checkout creates a new sales transaction and returns a sales receipt.
func Checkout(customerName string, items []*models.InventoryItem, paymentMethod string) (*SalesTransaction, *string, error) {
	// Calculate the total amount based on item prices and quantities
	var totalAmount float64
	for _, item := range items {
		totalAmount += item.Price * float64(item.Quantity)
	}

	// Generate a unique transaction ID (you can implement this logic)
	transactionID := generateUniqueTransactionID()

	// Create a SalesTransaction object
	transaction := &SalesTransaction{
		TransactionID:   transactionID,
		CustomerName:    customerName,
		TransactionDate: time.Now(),
		ItemsSold:       items,
		TotalAmount:     totalAmount,
		PaymentMethod:   paymentMethod,
	}

	// Save the transaction to the database (you can implement this logic)

	// Generate a sales receipt using the SalesReceipt.GenerateReceipt function
	receipt := GenerateReceipt(transaction)

	// Print or save the receipt as needed

	return transaction, &receipt, nil
}

// generateUniqueTransactionID generates a unique transaction ID.
func generateUniqueTransactionID() int {
	// Implement your logic to generate a unique ID (e.g., using a database sequence)
	// Return a unique ID here
	return 12345
}
