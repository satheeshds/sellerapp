package models

import "time"

type Inventory struct {
	Id         int       `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Price      float32   `json:"price,omitempty"`
	Status     string    `json:"status,omitempty"`
	Created_at time.Time `json:"created_at,omitempty"`
	Quantity   int       `json:"quantity,omitempty"`
}
