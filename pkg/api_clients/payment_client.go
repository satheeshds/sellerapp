package apiclients

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/satheeshds/sellerapp/pkg/models"
)

type PaymentClient struct{}

// This function captures a payment using the PaymentClient struct. It creates a new REST client and then sends a POST request to the payment processor with the payment information.
// If the response is successful, it returns nil, otherwise it returns an error containing the response string.
func (pc *PaymentClient) CapturePayment(payment *models.Payment) error {
	client := resty.New()

	resp, err := client.R().
		SetBody(payment).
		SetResult(payment).
		Post("http://payment_processor:5002/v1/payment")

	if err != nil {
		return err
	}

	if resp.IsSuccess() {
		return nil
	}

	return fmt.Errorf(resp.String())
}

// This function is part of the PaymentClient struct and is used to refund a payment.
// It takes in an integer paymentId as an argument.
// It creates a new resty client and uses it to send a PATCH request to the payment processor with the status set to "refunded".
// If the response is successful, it returns nil, otherwise it returns an error.
func (pc *PaymentClient) RefundPayment(paymentId int) error {
	client := resty.New()

	resp, err := client.R().
		SetBody(models.Payment{Status: "refunded"}).
		Patch(fmt.Sprintf("http://payment_processor:5002/v1/payment/%d", paymentId))

	if err != nil {
		return err
	}

	if resp.IsSuccess() {
		return nil
	}

	return fmt.Errorf(resp.String())
}
