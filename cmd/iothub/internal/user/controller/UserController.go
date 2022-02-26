package controller

import (
	"net/http"

	"github.com/lcnssantos/iothub/cmd/iothub/internal/user/dto"
	"github.com/lcnssantos/iothub/cmd/iothub/internal/user/service"
	http2 "github.com/lcnssantos/iothub/internal/http"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{userService: service}
}

func (this UserController) Create(w http.ResponseWriter, r *http.Request) {
	createUserDto := &dto.CreateUserDto{}

	if err := http2.HandleValidationRequest(w, r, createUserDto); err != nil {
		return
	}
	if err := this.userService.Create(*createUserDto, r.Context()); err != nil {
		http2.ThrowHttpError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	http2.SetResponse(w, 201, nil)
}

func (this UserController) List(w http.ResponseWriter, r *http.Request) {
	users, err := this.userService.List(r.Context())

	if err != nil {
		http2.ThrowHttpError(w, http.StatusInternalServerError, err.Error())
		return
	}

	http2.SetResponse(w, 200, users)
}
