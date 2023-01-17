package paymentrepository

import "github.com/satheeshds/sellerapp/pkg/models"

type IPaymentRepository interface {
	Create(*models.Payment) error
	Read(int) (models.Payment, error)
	Update(*models.Payment) error
	Open() error
	Close()
}
