package main

import (
	"net/http"

	"github.com/satheeshds/sellerapp/pkg/common"
	"github.com/satheeshds/sellerapp/pkg/models"
	paymentservice "github.com/satheeshds/sellerapp/pkg/payment_service"
)

// paymentapi is a struct that contains a field called paymentService which is of type IPaymentService from the paymentservice package.
// The paymentService field can be used to access methods and properties from the IPaymentService interface.
type paymentapi struct {
	paymentService paymentservice.IPaymentService
}

// This function is part of a payment API and is used to update a payment. It takes two parameters, an http.ResponseWriter and an http.Request, which are used to write the response to the request.
// The function first reads the body of the request into a payment object using the common.ReadBody() function. If there is an error reading the body, it writes an error response with a status code of 400 (Bad Request) and returns.
// If there is no error reading the body, it calls p.paymentService.UpdatePayment() to update the payment object with new information from the request body. If there is an error updating the payment, it writes an error response with a status code of 400 (Bad Request) and returns.
// If there is no error updating the payment, it writes a success response with a status code of 200 (OK) and returns with the updated payment object as part of its response body.
func (p *paymentapi) UpdatePayment(w http.ResponseWriter, r *http.Request) {
	var payment = &models.Payment{}
	if err := common.ReadBody(r, payment); err != nil {
		common.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	if err := p.paymentService.UpdatePayment(payment); err != nil {
		common.WriteErrorResponse(w, http.StatusBadRequest, err)
	} else {
		common.WriteSuccessResponse(w, http.StatusOK, payment)
	}
}

// This function is part of the paymentapi struct and is used to create a payment.
// It takes in a http.ResponseWriter and http.Request as parameters.
// It first reads the body of the request into a payment variable and then checks for any errors.
// If there are any errors, it will write an error response to the ResponseWriter with a status code of 400 (Bad Request).
// If there are no errors, it will call the CreatePayment function from the paymentService struct and check for any errors.
// If there are any errors, it will write an error response to the ResponseWriter with a status code of 500 (Internal Server Error).
// Otherwise, it will write a success response to the ResponseWriter with a status code of 201 (Created) and the payment variable as data.
func (p *paymentapi) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var payment = &models.Payment{}
	if err := common.ReadBody(r, payment); err != nil {
		common.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	if err := p.paymentService.CreatePayment(payment); err != nil {
		common.WriteErrorResponse(w, http.StatusInternalServerError, err)
	} else {
		common.WriteSuccessResponse(w, http.StatusCreated, payment)
	}
}

// This function is part of the paymentapi struct and is used to get a payment based on the given ID.
// It takes in two parameters, a http.ResponseWriter and an http.Request, and returns nothing.
// It first gets the ID from the request using mux.Vars(), then converts it to an integer using strconv.Atoi().
// If there is an error with this conversion, it writes a Bad Request error response using common.WriteErrorResponse().
// Otherwise, it calls p.paymentService.GetPayment() with the paymentid as a parameter and if there is an error, it writes a Not Found error response using common.WriteErrorResponse().
// Otherwise, it writes a Success response containing the result of p.paymentService.GetPayment() using common.WriteSuccessResponse().
func (p *paymentapi) GetPayment(w http.ResponseWriter, r *http.Request) {
	paymentid, err := common.ReadIdFromRequest(r)
	if err != nil {
		common.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	if result, err := p.paymentService.GetPayment(paymentid); err != nil {
		common.WriteErrorResponse(w, http.StatusNotFound, err)
	} else {
		common.WriteSuccessResponse(w, http.StatusOK, result)
	}
}
