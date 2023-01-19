package apiclients

import "testing"

func TestInventoryClient_BlockProduct(t *testing.T) {
	type args struct {
		product_id       int
		product_quantity int
	}
	tests := []struct {
		name    string
		i       *InventoryClient
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InventoryClient{}
			if err := i.BlockProduct(tt.args.product_id, tt.args.product_quantity); (err != nil) != tt.wantErr {
				t.Errorf("InventoryClient.BlockProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
