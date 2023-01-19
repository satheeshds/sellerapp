package models

import (
	"testing"
	"time"
)

func TestOrder_CancelOrder(t *testing.T) {
	type fields struct {
		Id               int
		User_id          int
		Transaction_id   int
		Status           string
		Created_at       time.Time
		Total_amount     float32
		Product_id       int
		Product_quantity int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Order{
				Id:               tt.fields.Id,
				User_id:          tt.fields.User_id,
				Transaction_id:   tt.fields.Transaction_id,
				Status:           tt.fields.Status,
				Created_at:       tt.fields.Created_at,
				Total_amount:     tt.fields.Total_amount,
				Product_id:       tt.fields.Product_id,
				Product_quantity: tt.fields.Product_quantity,
			}
			o.CancelOrder()
		})
	}
}
