package orderservice

import (
	"reflect"
	"testing"

	apiclients "github.com/satheeshds/sellerapp/pkg/api_clients"
	"github.com/satheeshds/sellerapp/pkg/models"
	orderrepository "github.com/satheeshds/sellerapp/pkg/order_repository"
)

func TestOrderService_GetOrder(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		o       *OrderService
		args    args
		want    models.Order
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "number",
			args: args{id: 2},
			want: models.Order{
				Id: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OrderService{}
			got, err := o.GetOrder(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderService.GetOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.GetOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderService_CreateOrder(t *testing.T) {
	type fields struct {
		Repository      orderrepository.IOrderRepository
		InventoryClient apiclients.IInventoryClient
		PaymentClient   apiclients.PaymentClient
	}
	type args struct {
		order *models.Order
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
			o := &OrderService{
				Repository:      tt.fields.Repository,
				InventoryClient: tt.fields.InventoryClient,
				PaymentClient:   tt.fields.PaymentClient,
			}
			if err := o.CreateOrder(tt.args.order); (err != nil) != tt.wantErr {
				t.Errorf("OrderService.CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrderService_CancelOrder(t *testing.T) {
	type fields struct {
		Repository      orderrepository.IOrderRepository
		InventoryClient apiclients.IInventoryClient
		PaymentClient   apiclients.PaymentClient
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Order
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OrderService{
				Repository:      tt.fields.Repository,
				InventoryClient: tt.fields.InventoryClient,
				PaymentClient:   tt.fields.PaymentClient,
			}
			got, err := o.CancelOrder(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderService.CancelOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderService.CancelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
