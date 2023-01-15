package orderservice

import (
	"reflect"
	"testing"

	"github.com/satheeshds/sellerapp/pkg/models"
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
