package orderservice

import (
	"github.com/satheeshds/sellerapp/pkg/models"
	orderrepository "github.com/satheeshds/sellerapp/pkg/order_repository"
)

type OrderService struct {
	Repository orderrepository.IOrderRepository
}

func (o *OrderService) GetOrder(id int) (models.Order, error) {
	return o.Repository.Read(id)
}

func (o *OrderService) CreateOrder(order models.Order) error {
	return o.Repository.Create(order)
}

func (o *OrderService) CancelOrder(id int) error {
	order, err := o.Repository.Read(id)
	if err != nil {
		return err
	}
	order.CancelOrder()
	return o.Repository.Update(id, order)
}
