package http

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func HandleValidationRequest(w http.ResponseWriter, r *http.Request, data interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		ThrowHttpError(w, http.StatusBadRequest, "Invalid Bod Request")
		return err
	}

	if err := validator.New().Struct(data); err != nil {
		ThrowHttpError(w, http.StatusBadRequest, err.Error())
		return err
	}

	return nil
}
