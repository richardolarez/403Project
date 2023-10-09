// models/inventory.go

package models

// InventoryItem represents an item in the pharmacy's inventory.
type InventoryItem struct {
	ID          int     // Unique identifier for the item
	Name        string  // Name of the item
	Description string  // Description of the item
	Price       float64 // Price of the item
	Quantity    int     // Quantity of the item in stock
}

// NewInventoryItem creates a new InventoryItem with the provided information.
func NewInventoryItem(id int, name string, description string, price float64, quantity int) *InventoryItem {
	return &InventoryItem{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}
}

// IncreaseQuantity increases the quantity of the item in stock.
func (item *InventoryItem) IncreaseQuantity(amount int) {
	item.Quantity += amount
}

// DecreaseQuantity decreases the quantity of the item in stock.
func (item *InventoryItem) DecreaseQuantity(amount int) {
	if item.Quantity >= amount {
		item.Quantity -= amount
	}
}

// TotalValue returns the total value of the item in stock (price * quantity).
func (item *InventoryItem) TotalValue() float64 {
	return item.Price * float64(item.Quantity)
}
