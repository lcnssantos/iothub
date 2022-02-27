package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func HandleValidationRequest(w http.ResponseWriter, r *http.Request, data interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		ThrowHttpException(w, http.StatusBadRequest, "Invalid Bod Request")
		return err
	}

	if err := validator.New().Struct(data); err != nil {
		ThrowHttpException(w, http.StatusBadRequest, err.Error())
		return err
	}

	return nil
}
