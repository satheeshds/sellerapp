package orderservice

import "github.com/satheeshds/sellerapp/pkg/models"

type IOrderService interface {
	GetOrder(id int) (models.Order, error)
	CreateOrder(*models.Order) error
	CancelOrder(id int) (models.Order, error)
}
