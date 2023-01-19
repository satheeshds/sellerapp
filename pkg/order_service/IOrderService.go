package orderservice

import "github.com/satheeshds/sellerapp/pkg/models"

// IOrderService is an interface that provides methods for interacting with orders.
// It has three methods:
// GetOrder() takes an integer id as an argument and returns a models.Order object and an error.
// CreateOrder() takes a pointer to a models.Order object as an argument and returns an error.
// CancelOrder() takes an integer id as an argument and returns a models.Order object and an error.
type IOrderService interface {
	GetOrder(id int) (models.Order, error)
	CreateOrder(*models.Order) error
	CancelOrder(id int) (models.Order, error)
}
