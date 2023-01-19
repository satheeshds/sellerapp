package paymentservice

import (
	"github.com/satheeshds/sellerapp/pkg/models"
	paymentrepository "github.com/satheeshds/sellerapp/pkg/payment_repository"
)

// PaymentService is a struct that contains a payment repository.
// It is used to provide access to payment related functions, such as creating and retrieving payments.
type PaymentService struct {
	Repository paymentrepository.IPaymentRepository
}

// This function is part of the PaymentService and is used to create a payment.
// It takes in a pointer to a Payment model as an argument and returns an error.
// The function calls the Create method on the Repository property of the PaymentService, passing in the payment argument.
func (p *PaymentService) CreatePayment(payment *models.Payment) error {
	return p.Repository.Create(payment)
}

// This function is part of the PaymentService and it is used to get a payment by its ID.
// It takes an integer as an argument and returns a models.Payment object as well as an error.
// The function calls the Repository's Read method to retrieve the payment from the database.
func (p *PaymentService) GetPayment(id int) (models.Payment, error) {
	return p.Repository.Read(id)
}

// This function is part of the PaymentService type and is used to update a payment.
// It takes a pointer to a Payment model as an argument and returns an error if one occurs.
// The function calls the Update method on the Repository field of the PaymentService type to update the payment.
func (p *PaymentService) UpdatePayment(payment *models.Payment) error {
	return p.Repository.Update(payment)
}
