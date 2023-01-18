package orderservice

import (
	apiclients "github.com/satheeshds/sellerapp/pkg/api_clients"
	"github.com/satheeshds/sellerapp/pkg/models"
	orderrepository "github.com/satheeshds/sellerapp/pkg/order_repository"
)

type OrderService struct {
	Repository      orderrepository.IOrderRepository
	InventoryClient apiclients.IInventoryClient
	PaymentClient   apiclients.PaymentClient
}

func (o *OrderService) GetOrder(id int) (models.Order, error) {
	return o.Repository.Read(id)
}

func (o *OrderService) CreateOrder(order *models.Order) error {
	if err := o.InventoryClient.BlockProduct(order.Product_id, order.Product_quantity); err != nil {
		return err
	}

	var payment = &models.Payment{
		Total_amount:   order.Total_amount,
		Payment_method: "cash",
	}

	if err := o.PaymentClient.CapturePayment(payment); err != nil {
		o.InventoryClient.UnblockProduct(order.Product_id, order.Product_quantity)
		return err
	}

	order.Transaction_id = payment.Id

	if err := o.Repository.Create(order); err != nil {
		o.InventoryClient.UnblockProduct(order.Product_id, order.Product_quantity)
		o.PaymentClient.RefundPayment(payment.Id)
		return err
	}

	return nil
}

func (o *OrderService) CancelOrder(id int) (models.Order, error) {
	order, err := o.Repository.Read(id)
	if err != nil {
		return order, err
	}
	o.InventoryClient.UnblockProduct(order.Product_id, order.Product_quantity)
	o.PaymentClient.RefundPayment(order.Transaction_id)
	order.CancelOrder()
	return order, o.Repository.Update(id, order)
}
