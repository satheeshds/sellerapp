package main

import (
	"net/http"
	"testing"

	inventoryservice "github.com/satheeshds/sellerapp/pkg/inventory_service"
)

func Test_inventoryapi_BlockInventory(t *testing.T) {
	type fields struct {
		inventoryService inventoryservice.IInventoryService
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
			i := &inventoryapi{
				inventoryService: tt.fields.inventoryService,
			}
			i.BlockInventory(tt.args.w, tt.args.r)
		})
	}
}

func Test_inventoryapi_CreateInventory(t *testing.T) {
	type fields struct {
		inventoryService inventoryservice.IInventoryService
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
			i := &inventoryapi{
				inventoryService: tt.fields.inventoryService,
			}
			i.CreateInventory(tt.args.w, tt.args.r)
		})
	}
}

func Test_inventoryapi_GetInventory(t *testing.T) {
	type fields struct {
		inventoryService inventoryservice.IInventoryService
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
			i := &inventoryapi{
				inventoryService: tt.fields.inventoryService,
			}
			i.GetInventory(tt.args.w, tt.args.r)
		})
	}
}
