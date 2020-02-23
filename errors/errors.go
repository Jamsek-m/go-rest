package errors

import "net/http"

// ************* Error response *************

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewErrorResponse(message string, status int) ErrorResponse {
	return ErrorResponse{
		Message: message,
		Status:  status,
	}
}

func NewErrorResponseFromError(error RestError) ErrorResponse {
	return ErrorResponse{
		Message: error.Message,
		Status:  error.Status,
	}
}

// ************* Base rest error *************

type RestError struct {
	Message string
	Status  int
}

func (err RestError) Error() string {
	return err.Message
}

func NewGenericRestError(message string) error {
	return RestError{Message: message, Status: http.StatusInternalServerError}
}

func NewRestError(message string, status int) error {
	return RestError{Message: message, Status: status}
}
