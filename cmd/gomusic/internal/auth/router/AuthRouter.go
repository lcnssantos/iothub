package router

import (
	"github.com/gorilla/mux"
	"github.com/lcnssantos/gomusic/cmd/gomusic/internal/auth/controller"
)

func BuildRouter(controller *controller.AuthController, router *mux.Router) {
	router.Methods("POST").HandlerFunc(controller.Auth)
}
