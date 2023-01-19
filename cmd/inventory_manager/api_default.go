package main

import (
	"net/http"

	"github.com/satheeshds/sellerapp/pkg/common"
	inventoryservice "github.com/satheeshds/sellerapp/pkg/inventory_service"
	"github.com/satheeshds/sellerapp/pkg/models"
)

// inventoryapi is a struct that contains an instance of the IInventoryService interface from the inventoryservice package.
// This struct can be used to access and manage inventory data.
type inventoryapi struct {
	inventoryService inventoryservice.IInventoryService
}

// This function is part of the inventoryapi and is used to block a certain amount of inventory.
// It takes in an http.ResponseWriter and an http.Request as parameters.
// First, it reads the body of the request and stores it in the variable 'inventory'. If there is an error, it writes an error response with a status code of 400 (Bad Request).
// Otherwise, it calls the BlockInventory() function from the inventoryService and if there is an error, it writes an error response with a status code of 400 (Bad Request).
// If there are no errors, it writes a success response with a status code of 200 (OK) and returns the 'inventory' variable.
func (i *inventoryapi) BlockInventory(w http.ResponseWriter, r *http.Request) {
	var inventory = &models.Inventory{}
	if err := common.ReadBody(r, inventory); err != nil {
		common.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err := i.inventoryService.BlockInventory(inventory); err != nil {
		common.WriteErrorResponse(w, http.StatusBadRequest, err)
	} else {
		common.WriteSuccessResponse(w, http.StatusOK, inventory)
	}
}

// This function creates an inventory item. It takes in a http.ResponseWriter and a http.Request as parameters and uses the readBody() function to read the body of the request.
// If there is an error while reading the body, it will write an error response with a status code of 400 (Bad Request).
// Otherwise, it will write a success response with a status code of 201 (Created) containing the inventory item.
func (i *inventoryapi) CreateInventory(w http.ResponseWriter, r *http.Request) {
	var inventory = &models.Inventory{}
	if err := common.ReadBody(r, inventory); err != nil {
		common.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err := i.inventoryService.CreateInventory(inventory); err != nil {
		common.WriteErrorResponse(w, http.StatusBadRequest, err)
	} else {
		common.WriteSuccessResponse(w, http.StatusCreated, inventory)
	}
}

// This function is part of the inventoryapi struct and is used to get an inventory item from the inventoryService.
// It takes in a http.ResponseWriter and a *http.Request as parameters.
// It then gets the id of the item from the request using mux.Vars and converts it to an integer using strconv.Atoi.
// If there is an error during this process, it will write an error response using common.WriteErrorResponse with a status code of http.StatusBadRequest and return from the function.
// Otherwise, it will attempt to get the inventory item from the inventoryService using GetInventory and if there is an error,
// it will write an error response with a status code of http.StatusNotFound, otherwise it will write a success response with a status code of http.StatusOK and pass in the result as a parameter.
func (i *inventoryapi) GetInventory(w http.ResponseWriter, r *http.Request) {
	productId, err := common.ReadIdFromRequest(r)
	if err != nil {
		common.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if result, err := i.inventoryService.GetInventory(productId); err != nil {
		common.WriteErrorResponse(w, http.StatusNotFound, err)
	} else {
		common.WriteSuccessResponse(w, http.StatusOK, result)
	}
}
