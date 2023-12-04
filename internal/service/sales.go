// service/sales.go

package service

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/richardolarez/403Project/internal/models"
)

var (
	lastID  int = 0
	idMutex sync.Mutex
)

// SalesTransaction represents a sales transaction with basic information and a list of items sold.
type Cart struct {
	CustomerID    string // Name of the customer associated with the transaction
	ID            string // Unique identifier for the transaction
	ItemID        string // Unique identifier for the item
	PaymentMethod string // Payment method used for the transaction
}

// Checkout creates a new sales transaction and returns a sales receipt.
func Checkout(customerID int, paymentMethod string, cartItems []*Cart) (*SalesReceipt, error) {
	items := []*models.InventoryItem{}

	for _, item := range cartItems {
		paymentMethod = item.PaymentMethod
		itemID, err := strconv.Atoi(item.ItemID)

		thisItem, err := models.GetInventoryItem(itemID)

		if err != nil {
			return nil, err
		}
		items = append(items, thisItem)
	}

	// Calculate the total amount based on item prices and quantities
	var totalAmount float64
	for _, item := range items {
		totalAmount += item.Price
	}

	// Generate a unique transaction ID (you can implement this logic)
	// transactionID := generateUniqueTransactionID()

	// // Create a SalesTransaction object
	// transaction := &models.SalesTransaction{
	// 	TransactionID:   transactionID,
	// 	CustomerID:      customerID,
	// 	TransactionDate: time.Now(),
	// 	ItemsSold:       items,
	// 	TotalAmount:     totalAmount,
	// 	PaymentMethod:   paymentMethod,
	// }

	// Retrieve or create the customer object using the customer repository
	// customer, err := models.GetCustomer(customerID)
	// if err != nil {
	// 	return "none", nil
	// }

	// // Append the transaction to the customer's transaction history
	// customer.AddTransaction(transaction)

	// Update the inventory quantities based on items sold
	if err := updateInventoryQuantities(items); err != nil {
		return nil, err
	}

	// Save the transaction and customer to the database (you can implement this logic)

	// Generate a sales receipt using the SalesReceipt.GenerateReceipt function
	receipt := GenerateReceipt(totalAmount, customerID, paymentMethod, items)

	// Print or save the receipt as needed

	return receipt, nil
}

// generateUniqueTransactionID generates a unique transaction ID.
func generateUniqueTransactionID() int {
	idMutex.Lock()
	lastID++
	idMutex.Unlock()
	return lastID
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
