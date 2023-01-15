package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	inventoryservice "github.com/satheeshds/sellerapp/pkg/inventory_service"
	"github.com/satheeshds/sellerapp/pkg/models"
)

type inventoryapi struct {
	inventoryService inventoryservice.IInventoryService
}

func (i *inventoryapi) BlockInventory(w http.ResponseWriter, r *http.Request) {
	inventory, err := readBody(r)
	if err != nil {
		log.Printf("error : %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	i.inventoryService.BlockInventory(inventory)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (i *inventoryapi) CreateInventory(w http.ResponseWriter, r *http.Request) {
	inventory, err := readBody(r)
	if err != nil {
		log.Printf("error : %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	i.inventoryService.CreateInventory(inventory)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

func (i *inventoryapi) GetInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	id := vars["id"]
	productId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := i.inventoryService.GetInventory(productId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	jsondata, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsondata)
}

func readBody(r *http.Request) (models.Inventory, error) {
	defer r.Body.Close()
	var inventory models.Inventory
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		err = json.Unmarshal(body, &inventory)
	}

	return inventory, err
}
