package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	orderservice "github.com/satheeshds/sellerapp/pkg/order_service"
)

func Test_orderapi_GetOrder(t *testing.T) {
	type fields struct {
		orderService orderservice.IOrderService
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
		{
			name: "first",
			fields: fields{
				orderService: &orderservice.OrderService{},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/1", nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &orderapi{
				orderService: tt.fields.orderService,
			}
			o.GetOrder(tt.args.w, tt.args.r)
		})
	}
}

func Test_orderapi_CancelOrder(t *testing.T) {
	type fields struct {
		orderService orderservice.IOrderService
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
			o := &orderapi{
				orderService: tt.fields.orderService,
			}
			o.CancelOrder(tt.args.w, tt.args.r)
		})
	}
}

func Test_orderapi_CreateOrder(t *testing.T) {
	type fields struct {
		orderService orderservice.IOrderService
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
			o := &orderapi{
				orderService: tt.fields.orderService,
			}
			o.CreateOrder(tt.args.w, tt.args.r)
		})
	}
}
