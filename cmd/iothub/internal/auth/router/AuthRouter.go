package router

import (
	"github.com/gorilla/mux"
	"github.com/lcnssantos/iothub/cmd/iothub/internal/auth/controller"
)

func BuildRouter(controller *controller.AuthController, router *mux.Router) {
	router.Methods("POST").HandlerFunc(controller.Auth)
	router.Methods("PUT").HandlerFunc(controller.Refresh)
}
