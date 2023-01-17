package main

import (
	"log"
	"net/http"
	"strings"

	apiclients "github.com/satheeshds/sellerapp/pkg/api_clients"
	"github.com/satheeshds/sellerapp/pkg/common"
	orderrepository "github.com/satheeshds/sellerapp/pkg/order_repository"
	orderservice "github.com/satheeshds/sellerapp/pkg/order_service"
)

func main() {

	log.Printf("Server starting!")

	orderRepository := &orderrepository.OrderRepository{}
	err := orderRepository.Open()
	if err != nil {
		log.Fatalf("Unable to open database connection: %v", err)
	}
	defer orderRepository.Close()

	orderService := orderservice.OrderService{
		Repository:      orderRepository,
		InventoryClient: &apiclients.InventoryClient{},
	}

	orderapi := &orderapi{
		orderService: &orderService,
	}
	var routes = common.Routes{

		common.Route{
			Name:        "CancelOrder",
			Method:      strings.ToUpper("Patch"),
			Pattern:     "/v1/order/{id}",
			HandlerFunc: orderapi.CancelOrder,
		},

		common.Route{
			Name:        "CreateOrder",
			Method:      strings.ToUpper("Post"),
			Pattern:     "/v1/order",
			HandlerFunc: orderapi.CreateOrder,
		},

		common.Route{
			Name:        "GetOrder",
			Method:      strings.ToUpper("Get"),
			Pattern:     "/v1/order/{id}",
			HandlerFunc: orderapi.GetOrder,
		},
	}
	router := common.NewRouter(routes)
	log.Printf("Server started!")
	log.Fatal(http.ListenAndServe(":5000", router))
}
