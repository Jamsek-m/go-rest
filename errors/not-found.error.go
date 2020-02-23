package errors

import "net/http"

type NotFoundError struct {
	Message string
}

func (err NotFoundError) Error() string {
	return err.Message
}

func NewNotFoundError(message string) error {
	return NewRestError(message, http.StatusNotFound)
}
