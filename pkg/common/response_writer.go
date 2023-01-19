package common

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/satheeshds/sellerapp/pkg/models"
)

func WriteErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(models.ErrorMessage{Message: err.Error()})
}

func WriteSuccessResponse(w http.ResponseWriter, statusCode int, model any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(model); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorMessage{Message: err.Error()})
	}
}

func ReadBody(r *http.Request, model any) error {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		err = json.Unmarshal(body, model)
	}
	return err
}
