package http

import (
	"net/http"
)

type HttpError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func ThrowHttpError(w http.ResponseWriter, status int, message string) {
	SetResponse(w, status, HttpError{Message: message, Status: status})
}
