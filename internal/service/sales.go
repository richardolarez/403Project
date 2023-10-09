// service/sales.go

package service

import (
	"fmt"
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
func Checkout(customerName string, items []*models.InventoryItem, paymentMethod string, customerRepo models.CustomerRepository) (*models.Customer, *SalesTransaction, error) {
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

	// Retrieve or create the customer object using the customer repository
	customer, err := customerRepo.GetByName(customerName)
	if err != nil {
		return nil, nil, err
	}

	// Append the transaction to the customer's transaction history using the customer repository
	if err := customerRepo.AddTransaction(customer, transaction); err != nil {
		return nil, nil, err
	}

	// Update the inventory quantities based on items sold
	if err := updateInventoryQuantities(items); err != nil {
		return nil, nil, err
	}

	// Save the transaction and customer to the database (you can implement this logic)

	// Generate a sales receipt using the SalesReceipt.GenerateReceipt function
	receipt := GenerateReceipt(transaction)

	// Print or save the receipt as needed

	return customer, transaction, nil
}

// generateUniqueTransactionID generates a unique transaction ID.
func generateUniqueTransactionID() int {
	// Implement your logic to generate a unique ID (e.g., using a database sequence)
	// Return a unique ID here
	return 12345
}

// updateInventoryQuantities updates the inventory item quantities after a sale.
func updateInventoryQuantities(items []*models.InventoryItem) error {
	// Implement logic to update the inventory quantities.
	// For example, you can decrement the quantity of each sold item.
	for _, item := range items {
		// Ensure that the quantity does not go negative
		if item.Quantity < 0 {
			return fmt.Errorf("item quantity cannot go negative: %s", item.Name)
		}

		// Decrement the quantity of the item
		item.DecreaseQuantity(1) // Decrement by 1 for each sold item
	}

	// Save the updated inventory quantities to the database (you can implement this logic)

	return nil
}
