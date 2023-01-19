package inventoryservice

import (
	"reflect"
	"testing"

	inventoryrepository "github.com/satheeshds/sellerapp/pkg/inventory_repository"
	"github.com/satheeshds/sellerapp/pkg/models"
)

func TestInventoryService_CreateInventory(t *testing.T) {
	type fields struct {
		Repository inventoryrepository.IInventoryRepository
	}
	type args struct {
		inventory models.Inventory
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
			i := &InventoryService{
				Repository: tt.fields.Repository,
			}
			if err := i.CreateInventory(&tt.args.inventory); (err != nil) != tt.wantErr {
				t.Errorf("InventoryService.CreateInventory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInventoryService_BlockInventory(t *testing.T) {
	type fields struct {
		Repository inventoryrepository.IInventoryRepository
	}
	type args struct {
		inventory models.Inventory
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
			i := &InventoryService{
				Repository: tt.fields.Repository,
			}
			if err := i.BlockInventory(&tt.args.inventory); (err != nil) != tt.wantErr {
				t.Errorf("InventoryService.BlockInventory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInventoryService_GetInventory(t *testing.T) {
	type fields struct {
		Repository inventoryrepository.IInventoryRepository
	}
	type args struct {
		pid int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Inventory
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InventoryService{
				Repository: tt.fields.Repository,
			}
			got, err := i.GetInventory(tt.args.pid)
			if (err != nil) != tt.wantErr {
				t.Errorf("InventoryService.GetInventory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InventoryService.GetInventory() = %v, want %v", got, tt.want)
			}
		})
	}
}
