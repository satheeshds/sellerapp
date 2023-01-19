package models

import "time"

// Payment is a struct that represents a payment made by a customer.
// It contains the following fields:
// Id: an integer representing the unique identifier of the payment
// Total_amount: a float32 representing the total amount of the payment
// Payment_method: a string representing the method used to make the payment (e.g. credit card, PayPal, etc.)
// Status: a string representing the status of the payment (e.g. successful, failed, pending, etc.)
// Created_at: a time.Time object representing when the payment was created
type Payment struct {
	Id             int       `json:"id,omitempty"`
	Total_amount   float32   `json:"total_amount,omitempty"`
	Payment_method string    `json:"payment_method,omitempty"`
	Status         string    `json:"status,omitempty" gorm:"default:successful"`
	Created_at     time.Time `json:"created_at,omitempty"`
}
