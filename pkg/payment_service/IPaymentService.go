package paymentservice

import "github.com/satheeshds/sellerapp/pkg/models"

// IPaymentService is an interface that provides methods for creating, getting, and updating payments.
// The CreatePayment method takes a pointer to a models.Payment struct as an argument and returns an error.
// The GetPayment method takes an integer as an argument and returns a models.Payment struct and an error.
// The UpdatePayment method takes a pointer to a models.Payment struct as an argument and returns an error.
type IPaymentService interface {
	CreatePayment(*models.Payment) error
	GetPayment(int) (models.Payment, error)
	UpdatePayment(*models.Payment) error
}
