package apiclients

import (
	"testing"

	"github.com/satheeshds/sellerapp/pkg/models"
)

func TestPaymentClient_CapturePayment(t *testing.T) {
	type args struct {
		payment *models.Payment
	}
	tests := []struct {
		name    string
		pc      *PaymentClient
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &PaymentClient{}
			if err := pc.CapturePayment(tt.args.payment); (err != nil) != tt.wantErr {
				t.Errorf("PaymentClient.CapturePayment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
