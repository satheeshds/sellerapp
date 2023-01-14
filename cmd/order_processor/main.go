package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/satheeshds/sellerapp/pkg/common"
)

func main() {
	log.Printf("Server started!")

	router := common.NewRouter(routes)

	log.Fatal(http.ListenAndServe(":5000", router))
}

var routes = common.Routes{

	common.Route{
		Name:        "CancelOrder",
		Method:      strings.ToUpper("Patch"),
		Pattern:     "/v1/order/{id}",
		HandlerFunc: CancelOrder,
	},

	common.Route{
		Name:        "CreateOrder",
		Method:      strings.ToUpper("Post"),
		Pattern:     "/v1/order",
		HandlerFunc: CreateOrder,
	},

	common.Route{
		Name:        "GetOrder",
		Method:      strings.ToUpper("Get"),
		Pattern:     "/v1/order/{id}",
		HandlerFunc: GetOrder,
	},
}
