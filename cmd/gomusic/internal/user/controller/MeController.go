package controller

import (
	http2 "github.com/lcnssantos/gomusic/internal/http"
	"net/http"
)

type MeController struct {
}

func NewMeController() *MeController {
	return &MeController{}
}

func (this MeController) GetMe(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	http2.SetResponse(w, http.StatusOK, user)
}
