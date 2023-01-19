package common

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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

// ReadIdFromRequest is a function that takes in a pointer to an http.Request and returns an int and an error.
// It uses the mux library to get the "id" variable from the request, then converts it to an integer using the strconv library's Atoi function.
func ReadIdFromRequest(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	id := vars["id"]
	return strconv.Atoi(id)
}
