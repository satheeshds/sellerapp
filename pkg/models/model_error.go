package models

// ErrorMessage is a struct used to store an error message in JSON format.
// It contains one field, "Message", which is a string and is set with the json tag "message,omitempty".
type ErrorMessage struct {
	Message string `json:"message,omitempty"`
}
