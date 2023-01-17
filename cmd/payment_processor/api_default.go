package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/satheeshds/sellerapp/pkg/models"
	paymentservice "github.com/satheeshds/sellerapp/pkg/payment_service"
)

type paymentapi struct {
	paymentService paymentservice.IPaymentService
}

func (p *paymentapi) UpdatePayment(w http.ResponseWriter, r *http.Request) {
	payment, err := readBody(r)
	if err != nil {
		log.Printf("error : %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorMessage{Message: err.Error()})
		return
	}
	err = p.paymentService.UpdatePayment(payment)
	if err != nil {
		log.Printf("error : %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorMessage{Message: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(payment); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorMessage{Message: err.Error()})
	}
}

func (p *paymentapi) CreatePayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	log.Printf("create payment")
	payment, err := readBody(r)
	if err != nil {
		log.Printf("error : %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorMessage{Message: err.Error()})
		return
	}
	if err := p.paymentService.CreatePayment(payment); err != nil {
		log.Printf("error : %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorMessage{Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(payment); err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{Message: err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (p *paymentapi) GetPayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	id := vars["id"]
	paymentid, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorMessage{Message: err.Error()})
		return
	}
	result, err := p.paymentService.GetPayment(paymentid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.ErrorMessage{Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{Message: err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func readBody(r *http.Request) (*models.Payment, error) {
	defer r.Body.Close()
	var inventory models.Payment
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		err = json.Unmarshal(body, &inventory)
	}

	return &inventory, err
}
