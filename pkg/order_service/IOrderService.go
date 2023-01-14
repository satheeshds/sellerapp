package orderservice

import "github.com/satheeshds/pkg/models"

type IOrderService interface {
	GetOrder(id int) models.Order
}
