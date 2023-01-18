package apiclients

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/satheeshds/sellerapp/pkg/models"
)

type PaymentClient struct{}

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
