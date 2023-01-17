package models

import "time"

type Order struct {
	Id               int       `json:"id,omitempty"`
	User_id          int       `json:"user_id,omitempty"`
	Transaction_id   int       `json:"transaction_id,omitempty"`
	Status           string    `json:"status,omitempty"`
	Created_at       time.Time `json:"created_at,omitempty"`
	Total_amount     float32   `json:"total_amount,omitempty"`
	Product_id       int       `json:"product_id,omitempty"`
	Product_quantity int       `json:"product_quantity,omitempty"`
}

func (o *Order) CancelOrder() {
	o.Status = "cancelled"
}
