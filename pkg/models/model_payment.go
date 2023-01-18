package models

import "time"

type Payment struct {
	Id             int       `json:"id,omitempty"`
	Total_amount   float32   `json:"total_amount,omitempty"`
	Payment_method string    `json:"payment_method,omitempty"`
	Status         string    `json:"status,omitempty" gorm:"default:successful"`
	Created_at     time.Time `json:"created_at,omitempty"`
}
