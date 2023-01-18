package apiclients

import "github.com/satheeshds/sellerapp/pkg/models"

type IPaymentClient interface {
	CapturePayment(*models.Payment) error
	RefundPayment(int) error
}
