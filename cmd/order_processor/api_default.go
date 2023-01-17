package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/satheeshds/sellerapp/pkg/models"
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
	defer r.Body.Close()
	var order models.Order
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error : %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal(body, &order)
	if err != nil {
		log.Fatalf("unmarshal -- %#v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = o.orderService.CreateOrder(order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

func (o *orderapi) GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	id := vars["id"]
	orderid, err := strconv.Atoi(id)
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
