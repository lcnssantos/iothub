package user

import "github.com/gorilla/mux"

func BuildRouter(controller *Controller, router *mux.Router) {
	router.Methods("POST").HandlerFunc(controller.Create)
}
