package common

import (
	"log"
	"net/http"
	"time"
)

// Logger is a function that takes in an http.Handler and a string as parameters and returns an http.Handler.
// The function creates a new http.HandlerFunc which logs the method, request URI, name, and time since start when it is called.
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
