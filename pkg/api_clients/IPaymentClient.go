package apiclients

import "github.com/satheeshds/sellerapp/pkg/models"

// IPaymentClient is an interface for interacting with a payment client.
// It contains two methods: CapturePayment and RefundPayment.
// The CapturePayment method takes in a pointer to a models.Payment struct and returns an error.
// The RefundPayment method takes in an integer and returns an error.
type IPaymentClient interface {
	CapturePayment(*models.Payment) error
	RefundPayment(int) error
}
