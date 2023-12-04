// service/sales_receipt.go

package service

import (
	"github.com/richardolarez/403Project/internal/models"
)

// SalesReceipt represents a sales receipt.
type SalesReceipt struct {
	TransactionID int
	Items         []*models.InventoryItem
	CustomerID    int
	TotalAmount   float64
	PaymentMethod string
	// Add more fields as needed
}

// GenerateReceipt generates a sales receipt object.
func GenerateReceipt(totalAmount float64, customerID int, paymentMethod string, items []*models.InventoryItem) *SalesReceipt {
	receipt := &SalesReceipt{
		TransactionID: 0, // Set the transaction ID as needed
		Items:         items,
		CustomerID:    customerID,
		TotalAmount:   totalAmount,
		PaymentMethod: paymentMethod,
		// Initialize other fields as needed
	}

	return receipt
}
