package orderservice

import (
	apiclients "github.com/satheeshds/sellerapp/pkg/api_clients"
	"github.com/satheeshds/sellerapp/pkg/models"
	orderrepository "github.com/satheeshds/sellerapp/pkg/order_repository"
)

// OrderService is a struct that contains the necessary components for handling orders.
// It contains a Repository field of type orderrepository.IOrderRepository, which is an interface for interacting with a repository of orders.
// It also contains two fields of type apiclients.IInventoryClient and apiclients.PaymentClient, which are interfaces for interacting with inventory and payment APIs respectively.
type OrderService struct {
	Repository      orderrepository.IOrderRepository
	InventoryClient apiclients.IInventoryClient
	PaymentClient   apiclients.PaymentClient
}

// This function is part of the OrderService and is used to get an order from the repository based on its ID.
// It takes in an integer representing the ID of the order and returns a models.Order object and an error.
func (o *OrderService) GetOrder(id int) (models.Order, error) {
	return o.Repository.Read(id)
}

// This function is part of the OrderService and is used to create an order. It takes in a pointer to a models.Order struct as an argument.
// First, it calls the BlockProduct function from the InventoryClient with the product ID and quantity from the order as arguments.
// If there is an error, it returns that error. Otherwise, it creates a payment struct with the total amount from the order and sets the payment method to "cash".
// It then calls CapturePayment from PaymentClient with the payment struct as an argument. If there is an error, it calls UnblockProduct on InventoryClient with the product ID and quantity from the order as arguments, then returns that error.
// If no errors occur, it sets Transaction_id in order to Payment's Id and calls Create on Repository with order as an argument.
// If there is an error, it calls UnblockProduct on InventoryClient with product ID and quantity from order as arguments, then calls RefundPayment on PaymentClient with Payment's Id as an argument before returning that error.
// If no errors occur during any of these steps, nil is returned.
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

// This function is part of the OrderService and is used to cancel an order. It takes in an integer id as a parameter and returns a models.Order and an error.
// It first reads the order from the Repository using the given id, and if there is an error, it returns that error.
// If there is no error, it then unblocks the product associated with the order using the InventoryClient and refunds the payment associated with the order using the PaymentClient.
// Finally, it calls CancelOrder() on the order object and updates it in the Repository before returning it along with any errors that may have occurred.
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
