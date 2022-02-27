package controller

import (
	"net/http"

	http2 "github.com/lcnssantos/iothub/internal/http"
)

type MeController struct {
}

func NewMeController() *MeController {
	return &MeController{}
}

func (c MeController) GetMe(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	http2.SendHttpResponse(w, http.StatusOK, user)
}
