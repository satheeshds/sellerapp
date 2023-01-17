package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/satheeshds/sellerapp/pkg/common"
	paymentrepository "github.com/satheeshds/sellerapp/pkg/payment_repository"
	paymentservice "github.com/satheeshds/sellerapp/pkg/payment_service"
)

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
