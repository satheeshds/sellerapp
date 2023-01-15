package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/satheeshds/sellerapp/pkg/common"
	inventoryrepository "github.com/satheeshds/sellerapp/pkg/inventory_repository"
	inventoryservice "github.com/satheeshds/sellerapp/pkg/inventory_service"
)

func main() {

	log.Printf("Server starting!")

	inventoryRepository := &inventoryrepository.InventoryRepository{}
	err := inventoryRepository.Open()
	if err != nil {
		log.Fatalf("Unable to open database connection: %v", err)
	}
	defer inventoryRepository.Close()

	inventoryService := inventoryservice.InventoryService{
		Repository: inventoryRepository,
	}

	api := &inventoryapi{
		inventoryService: &inventoryService,
	}
	var routes = common.Routes{

		common.Route{
			Name:        "createInventory",
			Method:      strings.ToUpper("Post"),
			Pattern:     "/v1/inventory",
			HandlerFunc: api.CreateInventory,
		},

		common.Route{
			Name:        "blockInventory",
			Method:      strings.ToUpper("Patch"),
			Pattern:     "/v1/inventory/{id}",
			HandlerFunc: api.BlockInventory,
		},
		common.Route{
			Name:        "GetOrder",
			Method:      strings.ToUpper("Get"),
			Pattern:     "/v1/inventory/{id}",
			HandlerFunc: api.GetInventory,
		},
	}
	router := common.NewRouter(routes)
	log.Printf("Server started!")
	log.Fatal(http.ListenAndServe(":5001", router))
}
