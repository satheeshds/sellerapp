package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	orderservice "github.com/satheeshds/sellerapp/pkg/order_service"
)

type orderapi struct {
	orderService orderservice.IOrderService
}

func (o *orderapi) CancelOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (o *orderapi) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (o *orderapi) GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("id = %v", id)
	orderid, err := strconv.Atoi(id)
	log.Printf(" after conv id = %v", id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := o.orderService.GetOrder(orderid)
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
