package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/satheeshds/sellerapp/pkg/common"
	"github.com/satheeshds/sellerapp/pkg/models"
	orderservice "github.com/satheeshds/sellerapp/pkg/order_service"
)

// orderapi is a struct that contains an IOrderService interface from the orderservice package.
// It is used to access methods related to order services.
type orderapi struct {
	orderService orderservice.IOrderService
}

// This function is part of the orderapi struct and is used to cancel an order.
// It takes in a http.ResponseWriter and http.Request as parameters.
// It uses mux to get the id from the request and converts it to an integer using strconv.Atoi().
// It then calls the CancelOrder() method from the orderService struct, passing in the orderid as a parameter.
// If there is an error, it calls writeErrorResponse() with the appropriate status code and error message, otherwise it calls writeSuccessResponse() with the order object.
func (o *orderapi) CancelOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	orderid, err := strconv.Atoi(id)
	if err != nil {
		common.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	order, err := o.orderService.CancelOrder(orderid)
	if err != nil {
		common.WriteErrorResponse(w, http.StatusInternalServerError, err)
	} else {
		common.WriteSuccessResponse(w, http.StatusOK, order)
	}
}

// This function is part of the orderapi type and creates an order.
// It takes in a http.ResponseWriter and http.Request as parameters.
// It reads the body and stores it in a variable of type models.Order.
// It then checks for any errors while reading the body, and if there are any, it writes an error response to the ResponseWriter with a status code of 400 (Bad Request).
// If no errors occur, it then attempts to unmarshal the body into JSON format and store it in the order variable.
// If there is an error during this process, it again writes an error response to the ResponseWriter with a status code of 400 (Bad Request).
// If no errors occur, it calls on the orderService to create an order using the order variable, and if successful, returns a success response with the order object included in it.
func (o *orderapi) CreateOrder(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var order models.Order
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		common.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = json.Unmarshal(body, &order); err != nil {
		common.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	if err = o.orderService.CreateOrder(&order); err != nil {
		common.WriteErrorResponse(w, http.StatusBadRequest, err)
	} else {
		common.WriteSuccessResponse(w, http.StatusCreated, order)
	}
}

// This function is part of the orderapi and is used to get an order based on its ID.
// It takes in a http.ResponseWriter and a *http.Request as parameters.
// It then uses mux to get the ID from the request, converts it to an integer, and passes it to the GetOrder function of the orderService. If there is an error, it writes an error response with a status code of either 400 (Bad Request) or 404 (Not Found).
// If there is no error, it writes a success response with the result.
func (o *orderapi) GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	orderid, err := strconv.Atoi(id)
	if err != nil {
		common.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if result, err := o.orderService.GetOrder(orderid); err != nil {
		common.WriteErrorResponse(w, http.StatusNotFound, err)
		return
	} else {
		common.WriteSuccessResponse(w, http.StatusOK, result)
	}
}
