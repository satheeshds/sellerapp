package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	id             int    `json:"id,omitempty"`
	user_id        int    `json:"user_id,omitempty"`
	transaction_id int    `json:"transaction_id,omitempty"`
	status         string `json:"status,omitempty"`
	created_at     string `json:"created_at,omitempty"`
}

func (o *Order) CancelOrder() {
	o.status = "cancelled"
}
