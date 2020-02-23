package errors

import "net/http"

type InternalServerError struct {
	Message string
}

func (err InternalServerError) Error() string {
	return err.Message
}

func NewInternalServerError(message string) error {
	return NewRestError(message, http.StatusInternalServerError)
}
