package common

import (
	"encoding/json"
	"io/ioutil"
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

// This function reads the body of an HTTP request and stores the data in a given model.
// It takes two parameters:
// 1) r *http.Request - a pointer to an http request object
// 2) model any - a variable of any type to store the read data in
// The function starts by deferring the closing of the request body, then reads all of the data from it using ioutil.ReadAll().
// If there is no error, it then uses json.Unmarshal() to convert the read data into JSON format and store it in the given model. Finally, it returns any errors that occurred during this process.
func ReadBody(r *http.Request, model any) error {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		err = json.Unmarshal(body, model)
	}
	return err
}
