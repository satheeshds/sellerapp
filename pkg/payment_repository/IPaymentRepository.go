package paymentrepository

import "github.com/satheeshds/sellerapp/pkg/models"

// IPaymentRepository is an interface that defines methods for interacting with a payment repository.
// The Create() method takes in a pointer to a models.Payment struct and returns an error if one occurs.
// The Read() method takes in an integer and returns a models.Payment struct and an error if one occurs.
// The Update() method takes in a pointer to a models.Payment struct and returns an error if one occurs.
// The Open() method opens the payment repository and returns an error if one occurs.
// The Close() method closes the payment repository and returns an error if one occurs.
type IPaymentRepository interface {
	Create(*models.Payment) error
	Read(int) (models.Payment, error)
	Update(*models.Payment) error
	Open() error
	Close()
}
