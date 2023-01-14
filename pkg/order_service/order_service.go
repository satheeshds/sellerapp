package orderservice

import "github.com/satheeshds/sellerapp/pkg/models"

type OrderService struct {
}

func (o *OrderService) GetOrder(id int) (models.Order, error) {
	return models.Order{}, nil
}

func (o *OrderService) CreateOrder(order models.Order) error {
	return nil
}

func (o *OrderService) CancelOrder(id int) error {
	return nil
}
