package main

import (
	"net/http"
	"testing"

	paymentservice "github.com/satheeshds/sellerapp/pkg/payment_service"
)

func Test_paymentapi_UpdatePayment(t *testing.T) {
	type fields struct {
		paymentService paymentservice.IPaymentService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &paymentapi{
				paymentService: tt.fields.paymentService,
			}
			p.UpdatePayment(tt.args.w, tt.args.r)
		})
	}
}

func Test_paymentapi_CreatePayment(t *testing.T) {
	type fields struct {
		paymentService paymentservice.IPaymentService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &paymentapi{
				paymentService: tt.fields.paymentService,
			}
			p.CreatePayment(tt.args.w, tt.args.r)
		})
	}
}

func Test_paymentapi_GetPayment(t *testing.T) {
	type fields struct {
		paymentService paymentservice.IPaymentService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &paymentapi{
				paymentService: tt.fields.paymentService,
			}
			p.GetPayment(tt.args.w, tt.args.r)
		})
	}
}
