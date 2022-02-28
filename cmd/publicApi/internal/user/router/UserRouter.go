package router

import (
	"github.com/gorilla/mux"
	"github.com/lcnssantos/iothub/cmd/publicApi/internal/auth/middleware"
	"github.com/lcnssantos/iothub/cmd/publicApi/internal/user/controller"
)

func BuildUserRouter(controller *controller.UserController, router *mux.Router, middleware *middleware.AuthenticationMiddleware) {
	router.Methods("POST").HandlerFunc(controller.Create)

	protectedRouter := router.PathPrefix("").Subrouter()
	protectedRouter.Use(middleware.Handler)
	protectedRouter.Methods("GET").HandlerFunc(controller.List)
}
