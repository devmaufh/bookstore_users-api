package errors

import "net/http"

type RestError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Error:      "bad_request",
	}
}
