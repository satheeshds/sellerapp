package common

import (
	"encoding/json"
	"net/http"

	"github.com/satheeshds/sellerapp/pkg/models"
)

// WriteErrorResponse is a function that takes in an http.ResponseWriter, an integer representing the status code, and an error as parameters.
// It sets the content type of the response to "application/json; charset=UTF-8" and writes the status code to the response.
// It then encodes a models.ErrorMessage containing the error message into JSON and writes it to the response.
func WriteErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(models.ErrorMessage{Message: err.Error()})
}

// This function writes a success response to an HTTP request.
// It takes in a http.ResponseWriter, an int representing the status code, and a model of any type as parameters.
// It sets the content type to "application/json; charset=UTF-8" and then writes the header with the given status code.
// It then encodes the model into JSON and writes it to the response writer.
// If there is an error encoding the model, it writes a 500 Internal Server Error status code and encodes an ErrorMessage object with the error message as its message field.
func WriteSuccessResponse(w http.ResponseWriter, statusCode int, model any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(model); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorMessage{Message: err.Error()})
	}
}
