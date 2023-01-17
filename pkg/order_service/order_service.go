package orderservice

import (
	apiclients "github.com/satheeshds/sellerapp/pkg/api_clients"
	"github.com/satheeshds/sellerapp/pkg/models"
	orderrepository "github.com/satheeshds/sellerapp/pkg/order_repository"
)

type OrderService struct {
	Repository      orderrepository.IOrderRepository
	InventoryClient apiclients.IInventoryClient
}

func (o *OrderService) GetOrder(id int) (models.Order, error) {
	return o.Repository.Read(id)
}

func (o *OrderService) CreateOrder(order models.Order) error {
	if err := o.InventoryClient.BlockProduct(order.Product_id, order.Product_quantity); err != nil {
		return err
	}

	if err := o.Repository.Create(order); err != nil {
		o.InventoryClient.UnblockProduct(order.Product_id, order.Product_quantity)
		return err
	}

	return nil
}

func (o *OrderService) CancelOrder(id int) error {
	order, err := o.Repository.Read(id)
	if err != nil {
		return err
	}
	order.CancelOrder()
	return o.Repository.Update(id, order)
}
