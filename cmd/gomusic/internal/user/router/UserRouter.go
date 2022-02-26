package router

import (
	"github.com/gorilla/mux"
	"github.com/lcnssantos/gomusic/cmd/gomusic/internal/user/controller"
)

func BuildUserRouter(controller *controller.UserController, router *mux.Router) {
	router.Methods("POST").HandlerFunc(controller.Create)
	router.Methods("GET").HandlerFunc(controller.List)
}
