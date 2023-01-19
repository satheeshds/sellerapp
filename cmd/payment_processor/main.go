package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/satheeshds/sellerapp/pkg/common"
	paymentrepository "github.com/satheeshds/sellerapp/pkg/payment_repository"
	paymentservice "github.com/satheeshds/sellerapp/pkg/payment_service"
)

// This code is the main function for a payment API.
// It starts by printing a log message indicating that the server is starting.
// It then creates an instance of the PaymentRepository and attempts to open a connection to the database.
// If an error occurs, it will be logged and the program will exit.
// The PaymentService is then initialized with the PaymentRepository instance and an instance of PaymentAPI is created with the PaymentService instance.
// Routes are then defined for creating payments, refunding payments, and getting payments.
// Finally, a router is created with these routes and the server starts listening on port 5002.
func main() {

	log.Printf("Server starting!")

	paymentRepository := paymentrepository.PaymentRepository{}
	err := paymentRepository.Open()
	if err != nil {
		log.Fatalf("Unable to open database connection: %v", err)
	}
	defer paymentRepository.Close()

	paymentService := paymentservice.PaymentService{
		Repository: &paymentRepository,
	}

	api := &paymentapi{
		paymentService: &paymentService,
	}
	var routes = common.Routes{

		common.Route{
			Name:        "createPayment",
			Method:      strings.ToUpper("Post"),
			Pattern:     "/v1/payment",
			HandlerFunc: api.CreatePayment,
		},

		common.Route{
			Name:        "refundPayment",
			Method:      strings.ToUpper("Patch"),
			Pattern:     "/v1/payment/{id}",
			HandlerFunc: api.UpdatePayment,
		},
		common.Route{
			Name:        "GetPayment",
			Method:      strings.ToUpper("Get"),
			Pattern:     "/v1/payment/{id}",
			HandlerFunc: api.GetPayment,
		},
	}
	router := common.NewRouter(routes)
	log.Printf("Server started!")
	log.Fatal(http.ListenAndServe(":5002", router))
}
