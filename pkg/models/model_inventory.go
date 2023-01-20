package models

import "time"

// Inventory is a struct that contains information about an item in an inventory.
// It has the following fields:
// Id: an integer that uniquely identifies the item in the inventory
// Name: a string containing the name of the item
// Price: a float32 representing the price of the item
// Status: a string indicating whether the item is available or not
// Created_at: a time.Time object representing when the item was added to the inventory
// Quantity: an integer indicating how many of this item are in stock
type Inventory struct {
	Id         int       `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Price      float32   `json:"price,omitempty"`
	Created_at time.Time `json:"created_at,omitempty"`
	Quantity   int       `json:"quantity,omitempty"`
}
