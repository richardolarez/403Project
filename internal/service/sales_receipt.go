// service/sales_receipt.go

package service

import (
	"fmt"

	"github.com/richardolarez/403Project/internal/models"
)

// SalesReceipt represents a sales receipt.
type SalesReceipt struct {
	TransactionID int
	Items         []*models.InventoryItem
	TotalAmount   float64
	PaymentMethod string
	// Add more fields as needed
}

// GenerateReceipt generates a formatted sales receipt as a string.
func GenerateReceipt(transaction *models.SalesTransaction) string {
	// Initialize the receipt text
	receipt := fmt.Sprintf("Transaction ID: %d\n", transaction.TransactionID)
	receipt += fmt.Sprintf("Date and Time: %s\n", transaction.TransactionDate.Format("2006-01-02 15:04:05"))
	receipt += "Items Purchased:\n"

	// Loop through the purchased items and list them on the receipt
	for _, item := range transaction.ItemsSold {
		receipt += fmt.Sprintf("  - %s (Quantity: %d)\n", item.Name, item.Quantity)
	}

	// Add total amount and payment method
	receipt += fmt.Sprintf("Total Amount: $%.2f\n", transaction.TotalAmount)
	receipt += fmt.Sprintf("Payment Method: %s\n", transaction.PaymentMethod)

	// Add more fields as needed

	return receipt
}
