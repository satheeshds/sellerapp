package common

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// This function creates a new router using the mux package.
// It takes in a Routes type as an argument and iterates through each route in the Routes type.
// For each route, it sets the handler to the HandlerFunc of that route, then wraps it with a Logger function.
// Finally, it sets the methods, path, name, and handler for that route and returns the router.
func NewRouter(routes Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router

}
