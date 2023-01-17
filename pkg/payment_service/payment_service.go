package paymentservice

import (
	"github.com/satheeshds/sellerapp/pkg/models"
	paymentrepository "github.com/satheeshds/sellerapp/pkg/payment_repository"
)

type PaymentService struct {
	Repository paymentrepository.IPaymentRepository
}

func (p *PaymentService) CreatePayment(payment *models.Payment) error {
	return p.Repository.Create(payment)
}

func (p *PaymentService) GetPayment(id int) (models.Payment, error) {
	return p.Repository.Read(id)
}

func (p *PaymentService) UpdatePayment(payment *models.Payment) error {
	return p.Repository.Update(payment)
}
