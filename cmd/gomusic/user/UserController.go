package user

import (
	"net/http"

	http2 "github.com/lcnssantos/gomusic/internal/http"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{service: service}
}

func (this Controller) Create(w http.ResponseWriter, r *http.Request) {
	createUserDto := &CreateUserDto{}

	if err := http2.HandleValidationRequest(w, r, createUserDto); err != nil {
		return
	}
	if err := this.service.Create(*createUserDto); err != nil {
		http2.ThrowHttpError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	http2.SetResponse(w, 201, nil)
}
