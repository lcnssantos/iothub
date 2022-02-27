package http

import (
	"net/http"
)

type HttpError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func ThrowHttpException(w http.ResponseWriter, status int, message string) {
	SendHttpResponse(w, status, HttpError{Message: message, Status: status})
}
