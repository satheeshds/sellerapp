package paymentservice

import "github.com/satheeshds/sellerapp/pkg/models"

type IPaymentService interface {
	CreatePayment(*models.Payment) error
	GetPayment(int) (models.Payment, error)
	UpdatePayment(*models.Payment) error
}
