package paymentservice

import (
	"reflect"
	"testing"

	"github.com/satheeshds/sellerapp/pkg/models"
	paymentrepository "github.com/satheeshds/sellerapp/pkg/payment_repository"
)

func TestPaymentService_CreatePayment(t *testing.T) {
	type fields struct {
		Repository paymentrepository.IPaymentRepository
	}
	type args struct {
		payment *models.Payment
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PaymentService{
				Repository: tt.fields.Repository,
			}
			if err := p.CreatePayment(tt.args.payment); (err != nil) != tt.wantErr {
				t.Errorf("PaymentService.CreatePayment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPaymentService_GetPayment(t *testing.T) {
	type fields struct {
		Repository paymentrepository.IPaymentRepository
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Payment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PaymentService{
				Repository: tt.fields.Repository,
			}
			got, err := p.GetPayment(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("PaymentService.GetPayment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentService.GetPayment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaymentService_UpdatePayment(t *testing.T) {
	type fields struct {
		Repository paymentrepository.IPaymentRepository
	}
	type args struct {
		payment *models.Payment
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PaymentService{
				Repository: tt.fields.Repository,
			}
			if err := p.UpdatePayment(tt.args.payment); (err != nil) != tt.wantErr {
				t.Errorf("PaymentService.UpdatePayment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
