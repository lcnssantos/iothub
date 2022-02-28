package controller

import (
	"net/http"

	"github.com/lcnssantos/iothub/cmd/publicApi/internal/user/dto"
	"github.com/lcnssantos/iothub/cmd/publicApi/internal/user/service"
	http2 "github.com/lcnssantos/iothub/internal/http"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{userService: service}
}

func (c UserController) Create(w http.ResponseWriter, r *http.Request) {
	createUserDto := &dto.CreateUserDto{}

	if err := http2.HandleValidationRequest(w, r, createUserDto); err != nil {
		return
	}
	if err := c.userService.Create(*createUserDto, r.Context()); err != nil {
		http2.ThrowHttpException(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	http2.SendHttpResponse(w, 201, nil)
}

func (c UserController) List(w http.ResponseWriter, r *http.Request) {
	users, err := c.userService.List(r.Context())

	if err != nil {
		http2.ThrowHttpException(w, http.StatusInternalServerError, err.Error())
		return
	}

	http2.SendHttpResponse(w, 200, users)
}
