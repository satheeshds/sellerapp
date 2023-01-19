package models

import "time"

// Order is a struct that contains information about an order.
// It has fields for the order's id, user id, transaction id, status, creation time, total amount, product id and product quantity.
// The fields are tagged with json and gorm to specify how they should be encoded and stored in a database.
type Order struct {
	Id               int       `json:"id,omitempty"`
	User_id          int       `json:"user_id,omitempty"`
	Transaction_id   int       `json:"transaction_id,omitempty"`
	Status           string    `json:"status,omitempty" gorm:"default:created"`
	Created_at       time.Time `json:"created_at,omitempty"`
	Total_amount     float32   `json:"total_amount,omitempty"`
	Product_id       int       `json:"product_id,omitempty"`
	Product_quantity int       `json:"product_quantity,omitempty"`
}

// This function sets the status of an Order object to "cancelled". The function takes a pointer to an Order object as an argument.
func (o *Order) CancelOrder() {
	o.Status = "cancelled"
}
