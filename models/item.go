package models

// Item represents a product in the inventory.
// It includes information like product name, stock count, and details.
type Item struct {
	// ID is the unique identifier for the item (assigned by the database)
	ID int `json:"id"`

	// ProductName is the human-readable name of the product
	ProductName string `json:"product_name" binding:"required"`

	// StockCount represents the quantity of the product currently in stock
	StockCount int `json:"stock_count" binding:"required"`

	// ItemDetails is a brief description or additional notes about the product
	ItemDetails string `json:"item_details"`
}
